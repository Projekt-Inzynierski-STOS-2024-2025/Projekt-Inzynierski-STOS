use std::{collections::HashMap, fmt::{Debug, Display}, fs, path::PathBuf, str::FromStr};

use crate::messages::{File, Files};

#[cfg(test)]
mod tests;

static DATA_DIRECTORY: &'static str = "./data/";

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

fn create_directory<T>(name: &T) -> Result<PathBuf, String>
where
    T: Display
{
    let mut save_dir = match PathBuf::from_str(DATA_DIRECTORY) {
        Ok(p) => p,
        Err(e) => return Err(format!("{e}")) 
    };
    save_dir.push(name.to_string());
    match fs::create_dir_all(save_dir.clone()) {
        Ok(_) => Ok(save_dir),
        Err(e) => Err(format!("{e}"))
    }
}

// Create an empty directory at path and load up tha files
pub fn save_files<T>(files: Vec<File>, dir_id: T) -> Result<T, String>
where 
    T: Display
{
    let dir = create_directory(&dir_id)?;
    for file in files {
        let mut file_path = dir.clone();
        file_path.push(file.name);
        let res = save_data(&file.data, file_path.clone());
        if res.is_err() {
            return Err(format!("Error while saving file: {:?}", file_path));
        }
    }
    Ok(dir_id)
}

// Read file contents from directory
pub fn load_file_contents<T>(paths: Vec<String>, dir_id: T) -> Result<Vec<File>, String> 
where 
    T: Display
{
    let dir = create_directory(&dir_id)?;
    // TODO - change it to prost generated struct
    let mut res: Vec<File>=  Vec::new();
    for file_name in paths {
        let mut file_path = dir.clone();
        file_path.push(&file_name);
        let contents = match get_data(file_path.clone()) {
            Ok(data) => data,
            Err(_) => return Err(format!("Error while reading from file: {:?}", file_path))
        };
        res.push(File{data: contents, name: file_name});
    }
    Ok(res)
}




