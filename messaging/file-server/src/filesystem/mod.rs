use std::{fs, path::PathBuf};

// TODO - better error forwarding
fn save_data(data: &Vec<u8>, destination: PathBuf) -> Result<(), String> {
    match fs::write(destination, data) {
        Ok(_) => Ok(()),
        Err(e) => Err(e.to_string())
    }
}

fn get_data(source: PathBuf) -> Result<Vec<u8>, String> {
    let data = fs::read(source);
    match data {
        Ok(data) => Ok(data),
        Err(e) => Err(e.to_string())
    }
}
