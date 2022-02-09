// https://stackoverflow.com/questions/36477623/given-an-xor-and-sum-of-two-numbers-how-to-find-the-number-of-pairs-that-satisf
//  010
//  111
// +---
// 1001
//
// + operator : 0+0 = 0, 0+1 = 1, 1+0 = 1, 1+1 = 0
// + operator is very similar to XOR operator
//
//    010
//    111
// XOR---
//    101
//
// only the difference is carry-in bit
// if the previous bits for a,b are 1, then we add it also
// This is exactly (a AND b) * 2
//
//    010
//    111
// AND---
//    010
//
//
// 101(a XOR b) + 010(a AND b) * 2 = 1001(a + b)
//
// we have (a+b) as s, (a AND b) as a, so that the formula are:
//
// a XOR b = s - 2 * a
//
// a and b should be positive, so
//
// (1) s - 2 * a >= 0
// (2) (s - 2 * a) & a == 0 ('cause (x XOR y) AND (x AND y) is 0)
//
// we just solve them, and if x,y are not satisfy above restrictions, return No, else Yes

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
