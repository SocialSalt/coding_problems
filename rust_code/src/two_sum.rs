use std::collections::HashMap;

pub struct Solution;
impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut map: HashMap<i32, usize> = HashMap::new();
        for (i, num) in nums.iter().enumerate() {
            if map.contains_key(num) {
                return vec![map[num] as i32, i as i32];
            }
            map.insert(target - num, i);
        }
        return vec![0,0];
    }
}

pub struct Cases;
impl Cases {
    pub fn all_cases() {
        assert!(Cases::case1(), "case 1 failed");
        assert!(Cases::case2(), "case 2 failed");
        println!("all test cases for two_sum passed");
    }

    fn case1() -> bool {
        let nums: Vec<i32> = vec![2,7,11,15];
        let target: i32 = 9;
        return Solution::two_sum(nums, target) == vec![0,1];
    }

    fn case2() -> bool {
        let nums: Vec<i32> = vec![3,2,4];
        let target: i32 = 6;
        return Solution::two_sum(nums, target) == vec![1,2];
    }
}
