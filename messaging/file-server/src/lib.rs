use lapin::message::Delivery;

pub mod file_service; 
pub mod messages {
    include!(concat!(env!("OUT_DIR"), "/stos.messages.rs"));
}
pub mod rabbit;

pub fn handle_message(delivery: Delivery) -> Result<(), String> {
    Ok(())
}

