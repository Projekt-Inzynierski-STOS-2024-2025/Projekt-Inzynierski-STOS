use std::{collections::HashMap, path::Path};
use super::*;

#[tokio::test]
async fn save_files_test() {
    let mut test_files: HashMap<String, Vec<u8>> = HashMap::new();
    test_files.insert(
        "first.txt".to_owned(),
        b"Hello from first file!".to_vec()
    );
    test_files.insert(
        "second.c".to_owned(),
        b"Hello from c!".to_vec()
    );

    let res = save_files(test_files, -2137);

    assert!(res.is_ok());
    assert!(Path::new("./data/-2137/first.txt").exists());
    assert!(Path::new("./data/-2137/second.c").exists());

    let _ = fs::remove_dir_all("./data/-2137");
}


#[tokio::test]
async fn read_files_test() {
    let mut test_files: HashMap<String, Vec<u8>> = HashMap::new();
    test_files.insert(
        "first.txt".to_owned(),
        b"Hello from first file!".to_vec()
    );
    test_files.insert(
        "second.c".to_owned(),
        b"Hello from c!".to_vec()
    );
    let paths = test_files.keys().map(|e| e.to_owned()).collect();
    save_files(test_files, -2138).unwrap();

    let res = load_file_contents(paths, -2138);

    assert!(res.is_ok());
    let res = res.unwrap();
    let first_content = res.get("first.txt").unwrap();
    assert_eq!(b"Hello from first file!".to_vec(), first_content.to_owned());


    let _ = fs::remove_dir_all("./data/-2138");
}
