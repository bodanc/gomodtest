package calc

// math.go is located inside the calc/ package to provide utility functions, such as Add();
func Add(numbers ...int) int {

	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum

}
