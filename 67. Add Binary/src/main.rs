pub struct Solution {}

impl Solution {
    pub fn add_binary(a: String, b: String) -> String {
        if a.len() < b.len() {
            return Solution::add_binary(b, a);
        }
        let mut i = 0;
        let mut result: Vec<char> = Vec::new();
        let mut left = 0;
        let a_bytes: Vec<u8> = a.as_bytes().iter().copied().rev().collect();
        let b_bytes: Vec<u8> = b.as_bytes().iter().copied().rev().collect();
        while i < a_bytes.len() {
            let curr = 
                if i < b_bytes.len() {
                    (a_bytes[i] - ('0' as u8)) + (b_bytes[i] - ('0' as u8)) + left
                } else {
                    (a_bytes[i] - ('0' as u8)) + left
                };
            match curr {
                0 => {
                    result.push('0');
                    left = 0;
                },
                1 => {
                    result.push('1');
                    left = 0;
                },
                2 => {
                    result.push('0');
                    left = 1;
                },
                _ => {
                    result.push('1');
                    left = 1;
                }
            }
            i += 1;
        }
        if left != 0 {
            result.push((left+('0' as u8)) as char);
        }
        return result.iter().rev().collect();
    }
}

fn main() {
    let a = String::from("1111");
    let b = String::from( "101");
    println!("{}", Solution::add_binary(a,b));
}
