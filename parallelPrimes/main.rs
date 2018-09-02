use std::sync::mpsc;
use std::thread;

fn main() {
    let (prime_sender, prime_receiver) = mpsc::channel();

    thread::spawn(move || {
        for k in 1..100000 {
            if is_prime(k) { prime_sender.send(k).unwrap(); }
        }
    });

    for received in prime_receiver {
        println!("Got {}", received);
    }
}

fn is_prime(n: i32) -> bool {
    match n {
        1 => false,
        2 => true,
        k if k > 2 => {
            for i in 2..(k / 2 + 1) {
                if k % i == 0 {
                    return false;
                }
            }
            true
        }
        _ => false
    }
}

#[test]
fn test_prime() {
    let is_two_prime = is_prime(2);
    let is_four_prime = is_prime(4);
    assert!(is_two_prime);
    assert!(!is_four_prime);
}
