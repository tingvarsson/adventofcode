use std::fs::File;
use std::io::{BufRead, BufReader, Result};
use std::path::Path;

fn main() -> Result<()> {
    let path = Path::new("day1/input");
    let file = File::open(&path)?;
    let mut numbers: Vec<i32> = vec![];
    for line in BufReader::new(file).lines() {
        numbers.push(line?.parse().unwrap())
    }

    let mut sum = 0;
    for n in &numbers {
        sum += n
    }
    println!("{}", sum);

    sum = 0;
    let mut i = 0;
    let mut knownsums: Vec<i32> = vec![];
    loop {
        sum += numbers[i % numbers.len()];
        match knownsums.binary_search(&sum) {
            Ok(_) => break,
            Err(pos) => knownsums.insert(pos, sum),
        }
        i += 1;
    }

    println!("{}", sum);
    Ok(())
}
