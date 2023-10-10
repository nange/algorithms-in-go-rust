package arraylist

func MoveZeroes(nums []int) {
	res := RemoveElement(nums, 0)
	for res < len(nums) {
		nums[res] = 0
		res++
	}
}
