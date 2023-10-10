use super::remove_element;

pub fn move_zeroes(nums: &mut Vec<i32>) {
    let mut res = remove_element(nums, 0) as usize;
    while res < nums.len() {
        nums[res] = 0;
        res += 1;
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_move_zeroes() {
        let mut nums = vec![0, 1, 0, 3, 12];
        move_zeroes(&mut nums);
        assert_eq!(nums, vec![1, 3, 12, 0, 0]);
    }
}
