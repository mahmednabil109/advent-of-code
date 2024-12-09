use std::{
    collections::{HashMap, HashSet},
    fs, io,
};

fn main() -> Result<(), io::Error> {
    let input = fs::read_to_string("./input1")?;
    let parts = input.split("\n\n").collect::<Vec<_>>();

    // part 1 naive
    let mut res1 = 0i32;
    let mut order: HashMap<i32, Vec<i32>> = HashMap::new();
    for line in parts[0].lines() {
        let parts = line
            .split("|")
            .map(|k| k.parse::<i32>().unwrap())
            .collect::<Vec<_>>();
        let (key, value) = (parts[0], parts[1]);
        let entry = order.entry(key).or_insert(Vec::new());
        entry.push(value);
    }

    for report in parts[1].lines() {
        let points = report
            .split(",")
            .map(|p| p.parse::<i32>().unwrap())
            .collect::<Vec<_>>();
        let mut valid = true;
        'check: for i in 0..points.len() {
            if !order.contains_key(&points[i]) {
                continue;
            }
            let o = order.get(&points[i]).unwrap();
            for j in 0..i {
                if o.contains(&points[j]) {
                    valid = false;
                    break 'check;
                }
            }
        }
        if valid {
            res1 += points[points.len() / 2];
        }
    }
    println!("Part 1: {}", res1);

    // part 2 (topological sort)
    let mut res2 = 0i32;
    for report in parts[1].lines() {
        let points = report
            .split(",")
            .map(|p| p.parse::<i32>().unwrap())
            .collect::<Vec<_>>();
        let index: HashMap<&i32, usize> = (&points)
            .into_iter()
            .enumerate()
            .map(|(i, n)| (n, i))
            .collect();

        // construct adj list
        let mut adj_list: Vec<HashSet<i32>> = Vec::new();
        for i in 0..points.len() - 1 {
            adj_list.push(
                (i + 1..points.len())
                    .map(|i| i as i32)
                    .collect::<HashSet<_>>(),
            )
        }
        adj_list.push(HashSet::new());
        for (i, p) in (&points).into_iter().enumerate() {
            if !order.contains_key(p) {
                continue;
            }
            for n in order.get(p).unwrap() {
                match index.get(n) {
                    Some(idx) => {
                        adj_list[i].insert(*idx as i32);
                    }
                    _ => continue,
                }
            }
        }
        let mut correct = true;
        loop {
            let mut visited = vec![false; points.len()];
            let mut path: Vec<i32> = Vec::new();
            let mut cyclic = false;
            for i in 0..points.len() {
                if visited[i] {
                    continue;
                }
                if detect_cycle(i as i32, &adj_list, &mut path, &mut visited) {
                    cyclic = true;
                    break;
                }
            }
            if cyclic {
                correct = false;
                // println!("{:?}", points);
                let st = (&path)
                    .into_iter()
                    .position(|r| r == path.last().unwrap())
                    .unwrap();
                let size = path.len() - st;
                // for i in &path[st..] {
                //     print!("{} ", points[*i as usize]);
                // }
                // println!("");
                for i in 0..size - 1 {
                    let (idx1, idx2) = (path[i + st], path[i + st + 1]);
                    let (p, nxt) = (points[idx1 as usize], points[idx2 as usize]);
                    // println!("{} -> {}", p, nxt);
                    // if one of the ordering rules, skip
                    if let Some(set) = order.get(&p) {
                        if set.contains(&nxt) {
                            continue;
                        }
                    }
                    // println!("chosed {} -> {}", p, nxt);
                    // remove this edge
                    adj_list[idx1 as usize].remove(&idx2);
                    break;
                }
            } else {
                break;
            }
        }
        if correct {
            continue;
        }
        // do topology sort
        let mut visited = vec![false; points.len()];
        let mut sorted: Vec<i32> = Vec::new();
        for i in 0..points.len() {
            if visited[i] {
                continue;
            }
            topological_sort(i as i32, &adj_list, &mut visited, &mut sorted);
        }
        //
        // for i in sorted.iter().rev() {
        //     print!("{} ", points[*i as usize]);
        // }
        // println!("");
        // for l in adj_list {
        //     println!("\t{:?}", l);
        // }
        let idx = sorted[sorted.len() / 2];
        // println!("-> {}", points[idx as usize]);
        res2 += points[idx as usize];
    }
    println!("Part 2: {}", res2);
    Ok(())
}

fn topological_sort(
    idx: i32,
    list: &Vec<HashSet<i32>>,
    visited: &mut Vec<bool>,
    stack: &mut Vec<i32>,
) {
    visited[idx as usize] = true;
    for nxt in &list[idx as usize] {
        if !visited[*nxt as usize] {
            topological_sort(*nxt, list, visited, stack);
        }
    }
    stack.push(idx);
}

fn detect_cycle(
    idx: i32,
    list: &Vec<HashSet<i32>>,
    path: &mut Vec<i32>,
    visited: &mut Vec<bool>,
) -> bool {
    path.push(idx);
    if visited[idx as usize] {
        return true;
    }
    visited[idx as usize] = true;
    for nxt in &list[idx as usize] {
        if !visited[*nxt as usize] && detect_cycle(*nxt, list, path, visited) {
            return true;
        } else if path.contains(nxt) {
            path.push(*nxt);
            return true;
        }
    }
    path.pop();
    false
}
