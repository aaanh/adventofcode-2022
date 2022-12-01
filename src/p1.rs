use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

pub fn p1() {
  let max:i32 = 0;
  let old_max:i32 = 0;

  if let Ok(lines) = read_lines("./src/inputs/p1.txt") {
    for line in lines {
      if let Ok(ip) = line {
        println!("{}", ip)
      }
    }
  } else {
    println!("Failed to load input file.")
  }
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where P: AsRef<Path>, {
  let file = File::open(filename)?;
  Ok(io::BufReader::new(file).lines())
}