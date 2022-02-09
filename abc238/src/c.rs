use proconio::input;
use std::cmp::min;

const M: u128 = 998244353;

// calc sum of [1, x)
fn f(x: u128) -> u128 {
    x * (x + 1) / 2
}

fn main() {
    input! {
        n: u128,
    }
    let keta = n.to_string().len();

    let mut ans = 0;
    let mut p10 = 10;
    for _ in 1..keta + 1 {
        // [l, r)
        let l = p10 / 10;
        let r = min(n, p10 - 1);
        ans += f(r - l + 1) % M;
        p10 *= 10;
    }
    ans %= M;
    println!("{}", ans);
}
