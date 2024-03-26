pub mod file_service; 
pub mod messages {
    include!(concat!(env!("OUT_DIR"), "/stos.messages.rs"));
}

