use crate::messages::File;

use self::{filesystem::save_files, redis_repo::RedisClient};

mod filesystem;
mod redis_repo;

#[derive(Clone, Debug)]
pub struct FileService {
    rc: RedisClient
}

impl FileService {
    pub async fn new() -> Result<Self, String> {
        let rc = match RedisClient::new().await {
            Ok(rc) => rc,
            Err(s) => return Err(s)
        };

        Ok(FileService{rc})
    }

    pub async fn store_files(&mut self, files: Vec<File>) -> Result<(), String> {    
        let id = match self.rc.store_paths(files.iter()
            .map(|s| s.name.clone())
            .collect()
        ).await {
            Ok(id) => id,
            Err(s) => return Err(s.to_string())
        };
        match save_files(files, id) {
            Ok(_) => Ok(()),
            Err(e) => Err(e)
        }
    }

    pub async fn get_files(&mut self, id: i64) -> Result<Vec<File>, String> {
        let paths = match self.rc.get_paths(id).await {
            Ok(p) => p,
            Err(e) => return Err(e.to_string())
        };
        filesystem::load_file_contents(paths, id)
    }
}

