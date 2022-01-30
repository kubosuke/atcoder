use proconio::input;
use std::collections::HashSet;

/*
連結グラフのコスト最大値を求めよ。
高さをもとに重み付き有効グラフに変換。
重みの最大値はDPで解く。
*/

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
        m: usize,
        h: [i64; n], // 0-index
        paths: [(usize, usize); m] // 1-index
    }
    let graph: Vec<HashSet<(usize, i64)>> = {
        let mut tmp: Vec<HashSet<(usize, i64)>> = vec![HashSet::new(); n];
        for path in paths.iter() {
            let x = path.0 - 1;
            let y = path.1 - 1;
            if h[x] > h[y] {
                tmp[x].insert((y, h[x] - h[y]));
                tmp[y].insert((x, 2 * (h[y] - h[x])));
            } else if h[x] < h[y] {
                tmp[y].insert((x, h[y] - h[x]));
                tmp[x].insert((y, 2 * (h[x] - h[y])));
            }
        }
        tmp
    };

    let mut dp = vec![-100000000000; n];
    dp[0] = 0;

    let mut ans = 0;
    for v in 0..90 * n {
        let v = v % n;
        for vv in &graph[v] {
            chmax!(dp[vv.0], dp[v] + vv.1);
            chmax!(ans, dp[vv.0]);
        }
    }
    println!("{}", ans);
}
