package arrays

func Sum(numbers []int) (sum int) {

	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbers ...[]int) []int {
	lenghtOfNumbers := len(numbers)
	sums := make([]int, lenghtOfNumbers)

	for i, numbers := range numbers {
		sums[i] = Sum(numbers)
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}

	}
	return sums
}
