package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	var sums []int
	for _, nums := range numbers {
		sums = append(sums, Sum(nums))
	}
	return sums
}

func SumAllTails(numbers ...[]int) []int {
	var sums []int
	for _, nums := range numbers {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(nums[1:]))
		}
	}
	return sums
}
