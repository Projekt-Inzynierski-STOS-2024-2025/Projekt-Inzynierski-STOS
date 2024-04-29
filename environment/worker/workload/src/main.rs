use std::time::Instant;

const DIMS: usize = 7500;

pub fn main() {
    println!("Running worker");
    let start = Instant::now();
    workload::run_compute(DIMS);
    let elapsed = start.elapsed();
    println!("Finished running worker in {:?} ms", elapsed.as_millis())

}
