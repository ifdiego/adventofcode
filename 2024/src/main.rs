mod days;

use std::{env, fs};

fn main() {
    let args: Vec<String> = env::args().collect();
    let day = args.get(1).expect("usage: cargo run -- <day>");
    let path = format!("inputs/day{day}.txt");
    let input = fs::read_to_string(&path)
        .unwrap_or_else(|_| panic!("failed to read input file for day {}", day));
    match day.as_str() {
        "01" => days::day01::run(&input),
        _ => eprintln!("day {} not yet implemented", day),
    }
}
