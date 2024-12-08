use std::{fs, io};

fn main() -> Result<(), io::Error> {
    let input = fs::read_to_string("./input1")?;

    // part 1
    let mut res1 = 0i32;
    let chars: Vec<char> = input.chars().collect();
    let mut i = 0;
    while i < chars.len() {
        let c = chars[i];
        if c == 'm' {
            if &input[i..i + 4] == "mul(" {
                let mut n1 = String::new();
                let mut n2 = String::new();
                let mut valid = true;
                for j in i + 4..chars.len() {
                    if chars[j] == ',' {
                        i = j;
                        break;
                    }
                    if chars[j] < '0' || chars[j] > '9' {
                        valid = false;
                        break;
                    }
                    n1.push(chars[j]);
                }
                if !valid {
                    i += 1;
                    continue;
                }
                for j in i + 1..chars.len() {
                    if chars[j] == ')' {
                        i = j;
                        break;
                    }
                    if chars[j] < '0' || chars[j] > '9' {
                        valid = false;
                        break;
                    }
                    n2.push(chars[j]);
                }
                if !valid {
                    i += 1;
                    continue;
                }
                let n1i32 = n1.parse::<i32>().unwrap();
                let n2i32 = n2.parse::<i32>().unwrap();
                res1 += n1i32 * n2i32;
            }
        }
        i += 1;
    }

    println!("Part 1: {}", res1);

    // part 2
    let mut res2 = 0i32;
    let chars: Vec<char> = input.chars().collect();
    let mut i = 0;
    'outer: while i < chars.len() {
        match chars[i] {
            'm' => 'case: {
                if &input[i..i + 4] == "mul(" {
                    let mut n1 = String::new();
                    let mut n2 = String::new();
                    let mut valid = true;
                    for j in i + 4..chars.len() {
                        if chars[j] == ',' {
                            i = j;
                            break;
                        }
                        if chars[j] < '0' || chars[j] > '9' {
                            valid = false;
                            break;
                        }
                        n1.push(chars[j]);
                    }
                    if !valid {
                       break 'case;
                    }
                    for j in i + 1..chars.len() {
                        if chars[j] == ')' {
                            i = j;
                            break;
                        }
                        if chars[j] < '0' || chars[j] > '9' {
                            valid = false;
                            break;
                        }
                        n2.push(chars[j]);
                    }
                    if !valid {
                       break 'case;
                    }
                    let n1i32 = n1.parse::<i32>().unwrap();
                    let n2i32 = n2.parse::<i32>().unwrap();

                    res2 += n1i32 * n2i32;
                }
            }
            'd' => {
                if &input[i..i+7] == "don't()" {
                    match input[i..].find("do()") {
                        Some(idx) => i += idx,
                        _ => break 'outer,
                    }
                } 
            }
            _ => {},
        }
        i += 1;
    }

    println!("Part 2: {}", res2);

    Ok(())
}
