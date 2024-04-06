use lapin::{message::Delivery, Channel, options::BasicPublishOptions, BasicProperties};

use crate::{messages::{Files, LogEvent, LogType}, file_service::FileService};

pub async fn handle_file_message(delivery: &Delivery) -> Result<i64, String> {
    let files: Files = delivery.data.clone().into();
    // Suboptimal since no connection pooling, might migrate to deadpool or sth later
    let mut fs = FileService::new_async().await?;
    fs.store_files(files.content).await
}

pub async fn handle_log_send(id: i64, channel: &Channel) -> lapin::Result<()> {
    let mut message = LogEvent::default();
    message.r#type = LogType::Info.into();
    message.time = chrono::offset::Local::now().to_rfc3339();
    message.content = format!("saved files under id: {id}");
    let serialized: Vec<u8> = message.into();
    channel.basic_publish("stos", "log", BasicPublishOptions::default(), &serialized, BasicProperties::default()).await?.await?;
    Ok(())
}

pub async fn handle_error_send(e: String, channel: &Channel) -> lapin::Result<()> {
    let mut message = LogEvent::default();
    message.r#type = LogType::Error.into();
    message.time = chrono::offset::Local::now().to_rfc3339();
    message.content = format!("FileServer error: {e}");
    let serialized: Vec<u8> = message.into();
    channel.basic_publish("stos", "log", BasicPublishOptions::default(), &serialized, BasicProperties::default()).await?.await?;
    Ok(())
}
