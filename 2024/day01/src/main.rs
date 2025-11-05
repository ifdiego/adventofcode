use std::collections::HashMap;
use std::fs;

fn main() {
    let path = format!("src/input.txt");
    let input = fs::read_to_string(&path).unwrap_or_else(|_| panic!("failed to read input file"));

    let mut left = Vec::new();
    let mut right = Vec::new();

    for line in input.lines() {
        let fields: Vec<&str> = line.split_whitespace().collect();
        if fields.len() != 2 {
            continue;
        }
        if let (Ok(l), Ok(r)) = (fields[0].parse::<i32>(), fields[1].parse::<i32>()) {
            left.push(l);
            right.push(r);
        }
    }

    // part one
    let mut left_sorted = left.clone();
    let mut right_sorted = right.clone();
    left_sorted.sort_unstable();
    right_sorted.sort_unstable();

    let part_one: i32 = left_sorted
        .iter()
        .zip(right_sorted.iter())
        .map(|(l, r)| (l - r).abs())
        .sum();
    println!("part one: {}", part_one);

    // part two
    let mut freq = HashMap::new();
    for r in &right {
        *freq.entry(r).or_insert(0) += 1;
    }

    let part_two: i32 = left.iter().map(|l| l * freq.get(l).unwrap_or(&0)).sum();
    println!("part two: {}", part_two);
}
