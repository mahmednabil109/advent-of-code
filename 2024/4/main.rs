use std::{fs, io};

fn main() -> Result<(), io::Error> {
    let input = fs::read_to_string("./input1")?;
    let directions = vec![
        (0, 1),
        (1, 0),
        (0, -1),
        (-1, 0),
        (1, 1),
        (-1, -1),
        (-1, 1),
        (1, -1),
    ];
    // part 1
    let mut res1 = 0i32;
    let grid: Vec<Vec<char>> = input.lines().map(|l| l.chars().collect()).collect();
    for i in 0..grid.len() {
        for j in 0..grid[i].len() {
            if grid[i][j] == 'X' {
                // println!("{} {}", i, j);
                // for d in &directions {
                //     let result = collect(&grid, (i as i32, j as i32), d);
                //     // println!("{:?} {}", d, result);
                //     if result.as_str() == "XMAS" {
                //         res1 += 1;
                //     }
                // }
                res1 += (&directions)
                    .into_iter()
                    .map(|d| collect(&grid, (i as i32, j as i32), &d))
                    .filter(|c| (*c).as_str() == "XMAS")
                    .collect::<Vec<_>>()
                    .len() as i32;
            }
        }
    }
    println!("Part 1: {}", res1);

    // part2
    let mut res2 = 0i32;
    for i in 1..grid.len() - 1 {
        for j in 1..grid[i].len() - 1 {
            if grid[i][j] == 'A' {
                let (mut x1, mut x2) = (String::new(), String::new());
                x1.push(grid[i - 1][j - 1]);
                x1.push(grid[i + 1][j + 1]);
                x2.push(grid[i - 1][j + 1]);
                x2.push(grid[i + 1][j - 1]);
                if (x1.as_str() == "MS" || x1.as_str() == "SM")
                    && (x2.as_str() == "MS" || x2.as_str() == "SM")
                {
                    res2 += 1;
                }
            }
        }
    }
    println!("Part 2: {}", res2);
    Ok(())
}

fn collect(grid: &Vec<Vec<char>>, start: (i32, i32), dir: &(i32, i32)) -> String {
    let mut counter = 0i32;
    let mut result = String::new();
    let (mut i, mut j) = start;

    while i >= 0
        && i < grid.len() as i32
        && j >= 0
        && j < grid[i as usize].len() as i32
        && counter < 4
    {
        // println!("   {} {}", i, j);
        result.push(grid[i as usize][j as usize]);
        i += dir.0;
        j += dir.1;
        counter += 1;
    }

    result
}
