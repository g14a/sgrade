use std::sync::mpsc;
use std::thread;

type FilterFn<T> = fn(T) -> bool;

fn main() {
    let (all_sender, all_receiver) = mpsc::channel();
    let (prime_sender, prime_receiver) = mpsc::channel();
    let (final_sender, final_receiver) = mpsc::channel();

    let is_prime: FilterFn<i32> = |num| {
        match num {
            1 => false,
            2 => true,
            k if k > 2 => {
                let sqrt =  ((k as f64).sqrt() as i64) + 1;
                for i in 2..sqrt {
                    if k as i64 % i == 0 {
                        return false;
                    }
                }
                true
            }
            _ => false
        }
    };

    let final_filter: FilterFn<i32> = |_num| { true };

    thread::spawn(move || {
        for k in 1..20000 {
            all_sender.send(k).unwrap();
        }
    });

    thread::spawn(move || {
        for received in all_receiver {
            let prime_sender_clone = prime_sender.clone();
            thread::spawn(move || {
                if is_prime(received) {
                    prime_sender_clone.send(received).unwrap();
                }
            });
        }
    });

    thread::spawn(move || {
        for received in prime_receiver {
            //println!("Got {}", received);
            let final_sender_clone = final_sender.clone();
            thread::spawn(move || {
                if final_filter(received) {
                    final_sender_clone.send(received).unwrap();
                }
            });
        }
    });

    for received in final_receiver {
        println!("Got {}", received);
    }

}

#[test]
fn test_prime() {
    let is_two_prime = is_prime(2);
    let is_four_prime = is_prime(4);
    assert!(is_two_prime);
    assert!(!is_four_prime);
}
