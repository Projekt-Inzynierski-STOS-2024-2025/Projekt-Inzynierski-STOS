#[tokio::main]
async fn main() {
    println!("Starting file server...");
    #[cfg(feature="benchmark")]
    {
        use file_server::file_service::FileService;
        use std::time::Instant;
        use file_server::messages::File;

        let RUNS = 100000;
        println!("Starting io benchmark.");
        dbg!(RUNS);
        let mut test_files: Vec<File> = Vec::new();
        test_files.push(File{
            name: "first.txt".to_owned(),
            data: b"Hello from first file!".to_vec()
        });
        test_files.push(File{
            name: "second.c".to_owned(),
            data: b"Hello from c!".to_vec()
        });

        println!("Starting single routine benchmark");
        let start = Instant::now();
        let mut fs = FileService::new().await.unwrap();
        for _ in 0..RUNS {
            let better_files = test_files.clone();
            let _ = fs.store_files(better_files).await.unwrap();
        }
        let time = start.elapsed();
        dbg!(time);
        println!("Finished benchmark. Running cleanup");
        let _ = std::fs::remove_dir_all("./data");
    }
    
}
