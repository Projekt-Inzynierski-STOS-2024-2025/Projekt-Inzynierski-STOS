use std::io::Cursor;

use lapin::message::Delivery;
use prost::{Message, DecodeError};

pub fn deserialize_message<T>(delivery: Delivery) -> Result<T, DecodeError>
where
    T: Message + Default
{
    let binary_content = delivery.data;
    let proto_struct = T::decode(&mut Cursor::new(binary_content))?;

    Ok(proto_struct)
}

pub fn serialize_message<T>(data: T) -> Vec<u8> 
where
    T: Message
{
    let mut buf = Vec::new();
    buf.reserve(data.encoded_len());
    data.encode(&mut buf).unwrap();
    buf
}
