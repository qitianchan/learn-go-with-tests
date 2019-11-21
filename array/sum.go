package main;

func Sum(numbers []int) int {
	sum := 0
	for _, n:= range numbers {
		sum += n
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	lenOfNumbers := len(numbersToSum)
	
	sums = make([]int, lenOfNumbers)
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return
}
