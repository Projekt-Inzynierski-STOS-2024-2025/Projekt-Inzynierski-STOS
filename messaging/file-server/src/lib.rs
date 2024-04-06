use messages::{Files, LogEvent};
use prost::Message;

pub mod file_service; 
pub mod messages {
    include!(concat!(env!("OUT_DIR"), "/stos.messages.rs"));
}
pub mod rabbit;

impl From<Vec<u8>> for Files {
    fn from(value: Vec<u8>) -> Self {
        match Files::decode(value.as_slice()) {
            Ok(files) => files,
            Err(e) => {
                eprintln!("{e}");
                Files::default()
            }
        }
    }
}

impl From<Files> for Vec<u8> {
    fn from(value: Files) -> Self {
        value.encode_to_vec()
    }
}

impl From<LogEvent> for Vec<u8> {
    fn from(value: LogEvent) -> Self {
        value.encode_to_vec()
    }
}
