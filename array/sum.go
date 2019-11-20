package main;

func Sum(numbers [5]int) (sum int) {
	sum = 0
	for i := range numbers {
		sum += numbers[i]
	}
	return
}