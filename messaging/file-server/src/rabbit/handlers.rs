use lapin::message::Delivery;

use crate::{messages::Files, file_service::FileService};

pub async fn handle_file_message(delivery: &Delivery) -> Result<i64, String> {
    let files: Files = delivery.data.clone().into();
    // Suboptimal since no connection pooling, might migrate to deadpool or sth later
    let mut fs = FileService::new_async().await?;
    fs.store_files(files.content).await
}
