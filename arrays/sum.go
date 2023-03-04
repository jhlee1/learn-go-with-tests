package arrays

func SumAll(numsOfNums ...[]int) []int {
	//result := make([]int, len(numsOfNums))
	var result []int

	for _, nums := range numsOfNums {
		result = append(result, Sum(nums))
	}

	return result
}

func Sum(nums []int) int {
	result := 0
	for _, num := range nums {
		result += num
	}

	return result
}

func SumAllTails(numsOfNums ...[]int) []int {
	var result []int
	for _, nums := range numsOfNums {
		if len(nums) < 1 {
			result = append(result, 0)
		} else {
			result = append(result, Sum(nums[1:]))
		}

	}

	return result
}
