pub fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {
    if nums.is_empty() {
        return 0;
    }

    let (mut slow, mut fast) = (0, 0);
    while fast < nums.len() {
        if nums[slow] != nums[fast] {
            slow += 1;
            nums[slow] = nums[fast]
        }
        fast += 1;
    }

    slow as i32 + 1
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_remove_duplicates() {
        let mut arr = vec![];
        let res = remove_duplicates(&mut arr);
        assert_eq!(res, 0);

        arr = vec![1, 2, 2, 3, 3, 3, 4, 4, 4, 4];
        let res = remove_duplicates(&mut arr);
        assert_eq!(res, 4);
        assert_eq!(arr[..4], vec![1, 2, 3, 4]);
    }
}
