use proconio::input;
use proconio::marker::Chars;

fn main() {
    input! {
        n: usize,
        s: Chars
    }
    let mut l: Vec<usize> = vec![];
    let mut r: Vec<usize> = vec![];

    for i in 0..n {
        if s[i] == 'L' {
            l.push(i);
        } else {
            r.push(i);
        }
    }
    r.push(n);

    for rr in r {
        print!("{} ", rr);
    }
    for ll in l.iter().rev() {
        print!("{} ", ll);
    }
}
