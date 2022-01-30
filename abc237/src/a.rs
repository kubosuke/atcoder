use proconio::input;

fn main() {
    input! {
        n: i64
    }
    if -2147483648 <= n && n <= 2147483647 {
        println!("Yes");
        return;
    }
    println!("No")
}
