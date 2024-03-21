# STOS file server

This is largely a POC solution. It utilizes redis as an async cache and saves files locally using filesystem.

## Running the server
File server requires redis to be installed locally and redis server to be running. You can run the server via `cargo run`

## Testing/benchmarking
There are unit tests as well as a simple benchmark testing the io operations. To test the application run `cargo test`. To run benchmarks, run the program with feature flag `benchmark` (`cargo run -F benchmark`)
