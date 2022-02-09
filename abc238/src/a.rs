use proconio::input;

fn main() {
    input! {
        n: f64,
    }
    println!("{}", {
        if n > 2.0 * n.log2() {
            "Yes"
        } else {
            "No"
        }
    });
}
