use proconio::input;
use proconio::marker::Chars;
use std::collections::VecDeque;

// corner: aaaaxaaxa

fn main() {
    input! {
        s: Chars
    }
    let mut dq: VecDeque<char> = VecDeque::from(s.to_vec());
    while !dq.is_empty() {
        let f = dq.pop_front();
        let b = dq.pop_back();

        if f == None || b == None {
            println!("Yes");
            return;
        }
        let f = f.unwrap();
        let b = b.unwrap();

        if f == b {
            continue;
        }
        if b == 'a' {
            dq.push_front(f);
            continue;
        }
        println!("No");
        return;
    }
    println!("Yes");
}
