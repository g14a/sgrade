const THRESHOLD: i32 = 10;


fn powerset<T>(s: &[T]) -> Vec<Vec<&T>> {
    (0..2usize.pow(s.len() as u32)).map(|i| {
         s.iter().enumerate().filter(|&(t, _)| (i >> t) % 2 == 1)
                             .map(|(_, element)| element)
                             .collect()
     }).collect()
}   

fn main() {

    let mut v = Vec::new();

    for k in 1..THRESHOLD {
        v.push(k);
    }

    println!("{:?}", v);
    let pset = powerset(&v);
    println!("{:?}", pset);
}
