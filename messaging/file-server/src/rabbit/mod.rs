use std::{sync::OnceLock, time::Duration};

use deadpool_lapin::{Pool, Manager, PoolError, Connection};
use futures::StreamExt;
use lapin::{ConnectionProperties, options::{QueueDeclareOptions, BasicConsumeOptions, BasicAckOptions}, types::FieldTable};
use tokio_amqp::LapinTokioExt;

use crate::rabbit::handlers::handle_file_message;

mod handlers;

pub fn pool() -> &'static Pool {
    static POOL_LOCK: OnceLock<Pool> = OnceLock::new();
    POOL_LOCK.get_or_init(|| {
        let addr = "amqp://guest:guest@127.0.0.1:5672/%2f".to_owned();
        let manager = Manager::new(addr, ConnectionProperties::default().with_tokio());
        let pool: Pool = deadpool::managed::Pool::builder(manager)
            .max_size(10)
            .build()
            .expect("can create pool");
        pool
    })
}

type RMQResult<T> = Result<T, PoolError>;

async fn get_rmq_con() -> RMQResult<Connection> {
    let connection = pool().get().await?;
    Ok(connection)
}

pub async fn rmq_listen() {
    let mut retry_interval = tokio::time::interval(Duration::from_secs(5));
    loop {
        retry_interval.tick().await;
        println!("connecting rmq consumer...");
        match init_rmq_listen().await {
            Ok(_) => println!("rmq listen returned"),
            Err(e) => eprintln!("rmq listen had an error: {:?}", e),
        };
    }
}

async fn init_rmq_listen() -> Result<(), lapin::Error> {
    let rmq_con = get_rmq_con().await.map_err(|e| {
        eprintln!("could not get rmq con: {}", e);
        e
    }).unwrap();
    let channel = rmq_con.create_channel().await.unwrap();

    let queue = channel
        .queue_declare(
            "files",
            QueueDeclareOptions::default(),
            FieldTable::default(),
        )
        .await.unwrap();
    println!("Declared queue {:?}", queue);

    let mut consumer = channel
        .basic_consume(
            "hello",
            "my_consumer",
            BasicConsumeOptions::default(),
            FieldTable::default(),
        )
        .await.unwrap();

    println!("rmq consumer connected, waiting for messages");
    while let Some(delivery) = consumer.next().await {
        if let Ok(delivery) = delivery {
            println!("received msg: {:?}", delivery);
            match handle_file_message(&delivery).await {
                Ok(id) => println!("Saved file with id: {id}"), // TODO - send log
                Err(e) => eprintln!("{e}") // TODO - send error log
            }
            channel
                .basic_ack(delivery.delivery_tag, BasicAckOptions::default())
                .await?
        }
    }
    Ok(())
}
