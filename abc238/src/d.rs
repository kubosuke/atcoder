use proconio::input;
use std::io::{self, Write};

fn main() {
    input! {
        t: usize,
    }
    let mut writer = io::BufWriter::new(io::stdout());
    for _ in 0..t {
        input! {
            a: i128,
            s: i128,
        }
        if s - 2 * a >= 0 && (s - 2 * a) & a == 0 {
            writeln!(writer, "Yes");
        } else {
            writeln!(writer, "No");
        }
    }
    writer.flush().unwrap();
}
