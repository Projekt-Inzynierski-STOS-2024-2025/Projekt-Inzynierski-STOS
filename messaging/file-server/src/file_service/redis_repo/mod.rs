use redis::{Client, AsyncCommands, aio::MultiplexedConnection, RedisResult};

#[cfg(test)]
mod tests;

// Default connection parameters in case env variables are not loaded
const REDIS_ADDRESS: &'static str = "127.0.0.1";
const COUNTER_KEY: &'static str = "path_counter";
const DB_KEY: &'static str = "paths";

#[derive(Clone, Debug)]
pub struct RedisClient {
    con: MultiplexedConnection
}

impl RedisClient {
    pub async fn new() -> Result<Self, String> {
        let address = REDIS_ADDRESS;
        let client_string = format!("redis://{address}/");
        let client = match Client::open(client_string) {
            Ok(c) => c,
            Err(e) => return Err(e.to_string())
        };
        let con = match client.get_multiplexed_async_connection().await {
            Ok(c) => c,
            Err(e) => return Err(e.to_string())
        };
        Ok(RedisClient{con})
    }

    // Automatically generate unique id utilizing redis counters
    async fn fetch_uuid(&mut self) -> RedisResult<i64> {
        self.con.incr(COUNTER_KEY, 1).await
    }

    // Save paths to local cache and return directory id
    pub async fn store_paths(&mut self, paths: Vec<String>) -> RedisResult<i64> {
        let id = match self.fetch_uuid().await {
            Ok(id) => id,
            Err(e) => return Err(e)
        };
        let key = format!("{DB_KEY}:{id}");
        let res: Result<(), redis::RedisError> = self.con.lpush(key, paths).await;
        match res {
            Ok(_) => Ok(id),
            Err(e) => Err(e)
        }

    }

    // Fetch paths saved under id
    pub async fn get_paths(&mut self, key: i64) -> RedisResult<Vec<String>>{
        let key = format!("{DB_KEY}:{key}");
        self.con.lrange(key, 0, -1).await
    }

}
