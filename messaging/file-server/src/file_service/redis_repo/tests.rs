use super::RedisClient;

#[tokio::test]
async fn test_connection() {
    let rc = RedisClient::new().await;

    assert!(rc.is_ok())
}

#[tokio::test]
async fn test_save() {
    let mut rc = RedisClient::new().await.unwrap();
    let test_paths = vec!["Fun path".to_owned(), "Good path".to_owned(), "Bad path".to_owned()];

    let id = rc.store_paths(test_paths).await;

    assert!(id.is_ok())
}


#[tokio::test]
async fn test_read() {
    let mut rc = RedisClient::new().await.unwrap();
    let test_paths = vec!["Fun path".to_owned(), "Good path".to_owned(), "Bad path".to_owned()];

    let id = rc.store_paths(test_paths.clone()).await.unwrap();
    let res = rc.get_paths(id).await.unwrap();

    // We need to reverse the iterator due to the way redis fetches lists
    assert!(res.iter().eq(test_paths.iter().rev()));
}


#[tokio::test]
async fn test_multiple_operations() {
    let mut rc = RedisClient::new().await.unwrap();
    let test_paths1 = vec!["Fun path".to_owned(), "Good path".to_owned(), "Bad path".to_owned()];
    let test_paths2 = vec!["Fun".to_owned(), "Good".to_owned(), "Bad".to_owned(), "Mediocre".to_owned()];

    let id1 = rc.store_paths(test_paths1.clone()).await.unwrap();
    let id2 = rc.store_paths(test_paths2.clone()).await.unwrap();
    //Check for unique ids
    assert_ne!(id1, id2);
    let res1 = rc.get_paths(id1).await.unwrap();
    let res2 = rc.get_paths(id2).await.unwrap();

    // We need to reverse the iterator due to the way redis fetches lists
    assert!(res1.iter().eq(test_paths1.iter().rev()));
    assert!(res2.iter().eq(test_paths2.iter().rev()));
}

