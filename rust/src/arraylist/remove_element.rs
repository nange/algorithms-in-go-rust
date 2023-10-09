pub fn remove_element(nums: &mut Vec<i32>, val: i32) -> i32 {
    let mut slow = 0;
    let mut fast = 0;

    while fast < nums.len() {
        if nums[fast] != val {
            if slow != fast {
                nums[slow] = nums[fast];
            }
            slow += 1;
        }
        fast += 1;
    }

    slow as i32
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_remove_element() {
        let mut nums = vec![3, 2, 2, 3];
        assert_eq!(remove_element(&mut nums, 3), 2);
        assert_eq!(remove_element(&mut nums, 4), 4);
    }
}
