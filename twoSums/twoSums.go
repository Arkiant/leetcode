package twosums

func twoSum(nums []int, target int) []int {
	result := make([]int, 2)
	values := make(map[int]int)
	for i, v := range nums {
		if val, ok := values[target-v]; ok {
			result[0] = val
			result[1] = i
			return result
		}
		values[v] = i
	}

	return result

}
