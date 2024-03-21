use std::{collections::HashMap, fmt::{Debug, Display}, fs, path::PathBuf, str::FromStr};

#[cfg(test)]
mod tests;

static DATA_DIRECTORY: &'static str = "./data/";

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
// TODO - protobuf + dto shared among services
pub fn save_files<T>(files: HashMap<String, Vec<u8>>, dir_id: T) -> Result<T, String>
where 
    T: Display
{
    let dir = match create_directory(&dir_id) {
        Ok(path) => path,
        Err(e) => return Err(e)
    };
    for (name, data) in files.iter() {
        let mut file_path = dir.clone();
        file_path.push(name);
        let res = save_data(data, file_path.clone());
        if res.is_err() {
            return Err(format!("Error while saving file: {:?}", file_path));
        }
    }
    Ok(dir_id)
}

// Read file contents from directory
pub fn load_file_contents<T>(paths: Vec<String>, dir_id: T) -> Result<HashMap<String, Vec<u8>>, String> 
where 
    T: Display
{
    let dir = match create_directory(&dir_id) {
        Ok(path) => path,
        Err(e) => return Err(e)
    };
    // TODO - change it to prost generated struct
    let mut res: HashMap<String, Vec<u8>>=  HashMap::new();
    for file_name in paths {
        let mut file_path = dir.clone();
        file_path.push(&file_name);
        let contents = match get_data(file_path.clone()) {
            Ok(data) => data,
            Err(_) => return Err(format!("Error while reading from file: {:?}", file_path))
        };
        res.insert(file_name, contents);
    }
    Ok(res)
}




