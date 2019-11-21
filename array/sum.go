package main;

func Sum(numbers []int) int {
	sum := 0
	for _, n:= range numbers {
		sum += n
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}
