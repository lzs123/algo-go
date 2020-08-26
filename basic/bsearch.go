package algo

func bSearch(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start + 1 < end {
		mid := start + (end - start) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			end = mid
		} else {
			start = mid
		}
	}
	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}
	return -1
}
