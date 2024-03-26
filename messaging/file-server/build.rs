use std::path::PathBuf;

extern crate prost_build;

fn main() {
    let current_path = std::env::current_dir().unwrap();
    let parent = current_path.parent().unwrap();
    let mut proto = PathBuf::from(parent);
    proto.push("messages.proto");
    let proto_files = [dbg!(proto.to_str().unwrap())];
    let dirs = [dbg!(parent.to_str().unwrap())];
    prost_build::compile_protos(&proto_files, &dirs).unwrap();
}
