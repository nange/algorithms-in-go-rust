package arraylist

func RemoveElement(nums []int, val int) int {
	slow, fast := 0, 0

	for fast < len(nums) {
		if nums[fast] != val {
			if slow != fast {
				nums[slow] = nums[fast]
			}
			slow++
		}
		fast++
	}

	return slow
}
