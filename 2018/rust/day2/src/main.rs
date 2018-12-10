use std::fs::File;
use std::collections::HashMap;
use std::io::{BufRead, BufReader, Result};
use std::path::Path;

fn main() -> Result<()> {
    let path = Path::new("../input");
    let file = File::open(&path)?;
    let mut lines: Vec<String> = vec![];
    for line in BufReader::new(file).lines() {
        lines.push(line?)
    }

    let mut sumoftwo = 0;
    let mut sumofthree = 0;
    for line in &lines {
        let mut letters: HashMap<char, i32> = HashMap::new();
        for c in line.chars() {
            if letters.contains_key(&c) {
                *letters.get_mut(&c).unwrap() += 1;
            } else {
                letters.insert(c, 1);
            }
        }

        let mut foundtwo = false;
        let mut foundthree = false;
        for (_key, value) in letters {
            if !foundtwo && value == 2 {
                foundtwo = true;
                sumoftwo += 1;
            } else if !foundthree && value == 3 {
                foundthree = true;
                sumofthree += 1;
            }
        }
    }
    println!("Checksum: {}", sumoftwo*sumofthree);

    for line in &lines {
        for secondline in &lines {
            let mut out: String = "".to_string();
            for (k, c) in line.chars().enumerate() {
                if c == secondline.chars().nth(k).unwrap() {
                    out += &c.to_string();
                }
            }
            if line.len()-1 == out.len() {
                println!("same line: {}", out);
                return Ok(())
            }
        }
    }

    Ok(())
}
