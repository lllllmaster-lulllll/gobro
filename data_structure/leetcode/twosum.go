package leetcode

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}
	numsMap := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		if _, ok := numsMap[nums[i]]; ok {
			return []int{numsMap[nums[i]], i}
		}
		numsMap[target-nums[i]] = i
	}
	return []int{}
}
