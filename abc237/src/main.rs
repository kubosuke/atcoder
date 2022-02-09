use proconio::input;
use proconio::marker::Usize1;
use std::cmp::Reverse;
use std::collections::BinaryHeap;

/*
連結グラフのコスト最大値を求めよ。
高さをもとに重み付き有効グラフに変換。
コストの最大値 = 負コストの最小値とみなす。
ポテンシャルを足しこんですべての経路を正にして、dijkstraで始点からのそれぞれの距離を求める。
そのなかの最小値からポテンシャルを引いて-1をかけあわせれば、本懐の最大コストが求まる。
*/

const INF: i64 = 2 << 60;

macro_rules! chmin {
    ($base:expr, $($cmps:expr),+ $(,)*) => {{
        let cmp_min = min!($($cmps),+);
        if $base > cmp_min {
            $base = cmp_min;
            true
        } else {
            false
        }
    }};
}

macro_rules! min {
    ($a:expr $(,)*) => {{
        $a
    }};
    ($a:expr, $b:expr $(,)*) => {{
        std::cmp::min($a, $b)
    }};
    ($a:expr, $($rest:expr),+ $(,)*) => {{
        std::cmp::min($a, min!($($rest),+))
    }};
}

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
    }
    let mut graph = vec![vec![]; 2 * m];
    for _ in 0..m {
        input! {
            mut u: Usize1,
            mut v: Usize1,
        };
        if h[u] > h[v] {
            graph[u].push((v, 0));
            graph[v].push((u, h[u] - h[v]));
        } else if h[u] < h[v] {
            graph[v].push((u, 0));
            graph[u].push((v, h[v] - h[u]));
        } else {
            graph[u].push((v, h[u] - h[v]));
            graph[v].push((u, h[v] - h[u]));
        }
    }

    // dijkstra
    let mut dist = vec![INF; n];
    let mut heap = BinaryHeap::new();
    heap.push((Reverse(0), 0));
    dist[0] = 0;

    while !heap.is_empty() {
        let (Reverse(d), v) = heap.pop().unwrap();
        if dist[v] != d {
            continue;
        };
        for &(to, cost) in &graph[v] {
            if chmin!(dist[to], dist[v] + cost) {
                heap.push((Reverse(dist[to]), to));
            }
        }
    }

    // seek
    let mut ans = 0;
    for i in 0..n {
        chmax!(ans, -(-h[0] + h[i] + dist[i]));
    }
    println!("{}", ans);
}
