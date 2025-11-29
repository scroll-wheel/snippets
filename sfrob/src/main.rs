use std::cmp::Ordering;
use std::io::{self, Read, Write};
use std::iter;

fn main() {
    let mut stdin = io::stdin();
    let mut buffer: Vec<u8> = Vec::new();
    stdin
        .read_to_end(&mut buffer)
        .expect("Failed to read from stdin");

    let space = 0x20;
    let mut lines = Vec::new();
    let mut line: Vec<u8> = Vec::new();
    for n in buffer {
        line.push(n);
        if n == space {
            lines.push(line.clone());
            line.clear();
        }
    }
    if !line.is_empty() {
        line.push(space);
        lines.push(line.clone());
    }

    lines.sort_unstable_by(|a, b| frobcmp(a, b));

    let mut stdout = io::stdout();
    for line in lines {
        stdout.write_all(&line).expect("Failed to read to stdout");
    }
}

fn frobcmp(a: &Vec<u8>, b: &Vec<u8>) -> Ordering {
    let newline = 0xA;
    let pattern = 0x2A;

    for (a, b) in iter::zip(a, b) {
        let a = a ^ pattern;
        let b = b ^ pattern;

        if a == newline && a == b {
            return Ordering::Equal;
        } else if a == newline || a < b {
            return Ordering::Less;
        } else if b == newline || a > b {
            return Ordering::Greater;
        }
    }

    assert!(false);
    Ordering::Equal
}
