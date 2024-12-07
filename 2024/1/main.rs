use std::{collections::HashMap, fs};

fn main(){
    let input = fs::read_to_string("./input1").expect("Error reading input file");

    // part 1
    let mut v1:Vec<i32> =  Vec::new();
    let mut v2:Vec<i32> =  Vec::new();

    for p in input.split("\n") {
        if p.len() == 0 {
            continue;
        }
        let idx = p.find(" ").unwrap();
        let p1 = p[0..idx].trim().parse::<i32>().unwrap();
        let p2 = p[idx..].trim().parse::<i32>().unwrap();
        v1.push(p1);
        v2.push(p2);
    }

    v1.sort();
    v2.sort();

    let mut res: i32 = 0;
    for i in 0..v1.len() {
        res += (v1[i] - v2[i]).abs()
    }
    
    println!("Part 1: {}", res);

    // part 2
    let mut freq: HashMap<i32, i32> = HashMap::new();
    for n in v2 {
        let e = freq.entry(n).or_insert(0);
        *e += 1;
    }

    let mut res2: i32 = 0;
    for n in v1 {
        let f = freq.get(&n).unwrap_or(&0);
        res2 += n * f;
    }

    println!("Part 2: {}", res2);

}
