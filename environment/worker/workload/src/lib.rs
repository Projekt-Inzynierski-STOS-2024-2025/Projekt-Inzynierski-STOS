use ndarray::Array;
use rand::Rng;

pub fn run_compute(dim: usize) {
    let mut a: Vec<f32> = Vec::new();
    let mut b: Vec<f32> = Vec::new();
    let mut rng = rand::thread_rng();

    for _ in 0..dim {
        for _ in 0..dim {
            a.push(get_random_val(&mut rng));
            b.push(get_random_val(&mut rng));
        }
    }
    let ma = Array::from_shape_vec((dim, dim), a).unwrap();
    let mb = Array::from_shape_vec((dim, dim), b).unwrap();
    let _ = ma.dot(&mb);
}

fn get_random_val(rng: &mut rand::rngs::ThreadRng) -> f32{
    1.4
}

