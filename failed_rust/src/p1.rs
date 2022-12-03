use std::fs::File;
use std::io::BufRead;
use std::io::BufReader;

pub fn p1() {
    let mut sum: u32 = 0;

    let file = File::open("./src/inputs/p1.txt").expect("File not found");
    let reader = BufReader::new(file);
    let mut numbers: Vec<u32> = Vec::new();

    for line in reader.lines() {
        if line.unwrap().parse::<u32>().is_err() {
            numbers.append(0);
        } else {
            numbers.append(line.unwrap().parse::<u32>());
        }
    }

    for num in numbers {
        sum += num;
    }

    println!("{}", sum);
}
