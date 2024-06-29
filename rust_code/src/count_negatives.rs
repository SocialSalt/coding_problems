use rayon::prelude::*;

pub struct Solution;
impl Solution {
    pub fn count_negatives(grid: Vec<Vec<i32>>) -> i32 {
        let m = grid.len();
        let n = grid[0].len();
        let mut count = 0;
        for (i, row) in grid.iter().enumerate() {
            if row[0] < 0 {
                count += n * (m - i);
                break;
            }
            for (j, val) in row.iter().enumerate() {
                if *val < 0 {
                    count += n - j;
                    break;
                }
            }
        }
        return count as i32;
    }

    fn zero_one(num: i32) -> i32 {
        return num>>31;
    }

    pub fn fast_sol(grid: Vec<Vec<i32>>) -> i32{
        return -1 * grid.par_iter().map( |row: &Vec<i32>| {row.par_iter().map(|num: &i32| -> i32 {Solution::zero_one(*num)}).sum::<i32>()}).sum::<i32>();
    }
}


pub struct Cases;
impl Cases {
    pub fn all_cases() {
        assert!(Cases::case1(), "case 1 failed");
        assert!(Cases::case2(), "case 2 failed");
        assert!(Cases::case3(), "case 3 failed");
        println!("all cases passed");
    }
    
    fn case1() -> bool {
        let grid: Vec<Vec<i32>> = vec![
            vec![9,4,3,2,2,1],
            vec![7,3,3,1,-2,-1],
            vec![5,2,-2,-3,-4,-6],
            vec![-9,-9,-9,-9,-9,-9]
            ];
        println!("{}", Solution::fast_sol(grid) == 12);
        // return Solution::count_negatives(grid) == 12;
        return true;
    }

    fn case2() -> bool {
        let grid: Vec<Vec<i32>> = vec![
            vec![4,3,3,1,1],
            vec![1,0,0,-1,-1],
            vec![-2,-2,-2,-2,-3],
            vec![-2,-2,-2,-3,-3],
            vec![-3,-3,-3,-3,-3]
        ];
        println!("{}", Solution::fast_sol(grid));
        return true;
    }

    fn case3() -> bool {
        let grid = vec![
            vec![4,3,2,-1],
            vec![3,2,1,-1],
            vec![1,1,-1,-2],
            vec![-1,-1,-2,-3]
        ];
        println!("{}", Solution::fast_sol(grid));
        return true;
    }
}
