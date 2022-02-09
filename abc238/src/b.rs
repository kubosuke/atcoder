use proconio::input;

macro_rules! chmax {
    ($base:expr, $($cmps:expr),+ $(,)*) => {{
        let cmp_max = max!($($cmps),+);
        if $base < cmp_max {
            $base = cmp_max;
            true
        } else {
            false
        }
    }};
}

macro_rules! max {
    ($a:expr $(,)*) => {{
        $a
    }};
    ($a:expr, $b:expr $(,)*) => {{
        std::cmp::max($a, $b)
    }};
    ($a:expr, $($rest:expr),+ $(,)*) => {{
        std::cmp::max($a, max!($($rest),+))
    }};
}

fn main() {
    input! {
        n: usize,
        a: [usize;n],
    }
    let mut v = vec![];
    v.push(a[0]);
    for i in 1..n {
        v = v
            .iter()
            .map(|x| {
                if x + a[i] > 360 {
                    x + a[i] - 360
                } else {
                    x + a[i]
                }
            })
            .collect();
        v.push(a[i]);
    }
    v.push(0);
    v.push(360);

    v.sort();

    let mut ans = 0;
    for i in 0..n + 1 {
        let diff = v[i + 1] - v[i];
        chmax!(ans, diff);
    }
    println!("{}", ans);
}
