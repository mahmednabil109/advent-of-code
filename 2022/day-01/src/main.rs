use std::fs;
use priority_queue::PriorityQueue;

fn main() {
    let input = fs::read_to_string("./input1")
                .expect("Error reading the file");

    // first half
    let mut elfs = input.split("\n\n");
    let mut max_cal = 0i32;
    for elf in elfs {
        let mut cal_sum = 0i32;
        for cal in elf.lines() {
            cal_sum += cal.parse::<i32>().unwrap();
        }
        if cal_sum > max_cal {
            max_cal = cal_sum;
        }
    }
    println!("{}", max_cal);

    // second half
    // TODO needs to be searched
    let elfs2 = input.split("\n\n");
    let mut pq = PriorityQueue::new();
    for elf in elfs2 {
        let mut cal_sum = 0i32;
        for cal in elf.lines() {
            cal_sum += cal.parse::<i32>().unwrap();
        }
        pq.push(cal_sum, cal_sum * -1);
        if pq.len() > 3 {
            pq.pop();
        }
    }

    assert!(pq.len() <= 3);
    let mut ans2 = 0i32;
    for (item, _) in pq.iter(){
        ans2 += item;
    }
    println!("{}", ans2);
}
