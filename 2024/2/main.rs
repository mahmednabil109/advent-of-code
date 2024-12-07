use std::{fs, io};

fn main() -> Result<(), io::Error> {
    let input = fs::read_to_string("./input1")?;

    // part 1
    let mut res = 0i32;
    for line in input.lines() {
        if line.len() == 0 {
            continue;
        }
        let levels: Vec<i32> = line.split(" ").map(|l| l.parse::<i32>().unwrap()).collect();

        if check_report(&levels) {
            res += 1;
        }
    }
    println!("Part 1: {}", res);

    // part 2
    let mut res2 = 0i32;
    for line in input.lines() {
        if line.len() == 0 {
            continue;
        }
        let levels: Vec<i32> = line.split(" ").map(|l| l.parse::<i32>().unwrap()).collect();

        if check_report(&levels) || tolerable(&levels) {
            res2 += 1;
        }
    }
    println!("Part 2: {}", res2);

    Ok(())
}

fn tolerable(levels: &Vec<i32>) -> bool {
    for i in 0..levels.len() {
        // This is so ugly
        let mut cloned = levels.clone();
        cloned.remove(i);
        if check_report(&cloned) {
            return true;
        }
    }
    false
}

fn check_report(levels: &Vec<i32>) -> bool {
    let inc: bool = levels[1] > levels[0];
    let mut safe: bool = true;
    for i in 0..levels.len() - 1 {
        if inc {
            if levels[i + 1] <= levels[i] || levels[i + 1] - levels[i] > 3 {
                safe = false;
                break;
            }
        } else {
            if levels[i + 1] >= levels[i] || levels[i] - levels[i + 1] > 3 {
                safe = false;
                break;
            }
        }
    }
    safe
}
