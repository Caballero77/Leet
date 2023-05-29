pub struct Solution {}

impl Solution {
    pub fn remove_element(nums: &mut Vec<i32>, val: i32) -> i32 {
        let mut count = 0;
        let mut i = 0;
        let len = nums.len();
        while i < (len-count) {
            if nums[i] == val {
                nums[len-1-count] += nums[i];
                nums[i] = nums[nums.len()-1-count]-nums[i];
                nums[len-1-count] -= nums[i];
                count += 1;
            } else {
                i += 1;
            }
        }
        return (len-count) as i32;
    }
}

fn main() {
    let mut vector = vec![0,1,2,2,3,0,4,2];
    println!("{}", Solution::remove_element(&mut vector, 2));
    for i in vector {
        print!("{}", i)
    }
}
