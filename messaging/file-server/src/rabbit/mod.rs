use std::ops::{Deref, DerefMut};
use deadpool::managed::{Manager, Pool};
use lapin::ConnectionProperties;

#[derive(Clone, Debug)]
pub struct RabbitClient {
    pool: Pool
}


impl Deref for RabbitClient {

    type Target = Pool<>;

    fn deref(&self) -> &Self::Target {
        &self.pool    
    }
}


impl DerefMut for RabbitClient {
    fn deref_mut(&mut self) -> &mut Self::Target {
        &mut self.pool
    }
}

impl RabbitClient  {
    pub fn new(pool_size: usize) -> Self {
        let addr = std::env::var("AMQP_ADDR")
            .unwrap_or_else(|_| "amqp://rmq:rmq@127.0.0.1:5672/%2f".into());
        let manager = Manager::new(addr, ConnectionProperties::default().with_tokio());
        let pool: Pool = deadpool::managed::Pool::builder(manager)
            .max_size(pool_size)
            .build()
            .expect("can create pool");

        RabbitClient { pool }
    } 
}
