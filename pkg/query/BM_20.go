package query

func reversePairsMain(nums []int) int {
	len := len(nums)
	if len < 2 {
		return 0
	}
	temp := make([]int, len)
	copy := nums
	return reversePairs(copy, 0, len-1, temp)
}

func reversePairs(nums []int, left, right int, temp []int) int {
	if left == right {
		return 0
	}
	mid := left + (right-left)>>1
	leftPairs := reversePairs(nums, left, mid, temp)
	rightPairs := reversePairs(nums, mid+1, right, temp)
	crossPairs := mergeAndCount(nums, left, mid, right, temp)
	return leftPairs + rightPairs + crossPairs
}

func mergeAndCount(nums []int, left, mid, right int, temp []int) int {
	for i := left; i <= right; i++ {
		temp[i] = nums[i]
	}
	i, j, count := left, mid+1, 0
	for k := left; k <= right; k++ {
		if i == mid+1 {
			nums[k] = temp[j]
			j++
		} else if j == right+1 {
			nums[k] = temp[i]
			i++
		} else if temp[i] <= temp[j] {
			nums[k] = temp[i]
			i++
		} else {
			nums[k] = temp[j]
			count += (mid + 1 - i)
			j++
		}
	}
	return count
}
