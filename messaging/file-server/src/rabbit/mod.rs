use std::ops::{Deref, DerefMut};
use deadpool_lapin::{Config, Pool};

pub mod messages;

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
    pub fn new() -> Self {
        let mut cfg = Config::default();
        cfg.url = Some("amqp://rmq:rmq@127.0.0.1:5672/%2f".into());
        let pool = cfg.create_pool(Some(deadpool::Runtime::Tokio1)).unwrap();
        RabbitClient { pool }
    } 

    pub fn run_messaging(&mut self) -> Result<(), String> {
        

        Ok(())
    }
}
