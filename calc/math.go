package calc

import "errors"

// math.go is located inside the calc/ package to provide utility functions, such as Add();
func Add(numbers ...int) (error, int) {

	sum := 0

	if len(numbers) < 2 {
		return errors.New("you need to provide more than 2 numbers"), sum
	} else {
		for _, number := range numbers {
			sum += number
		}

		return nil, sum
	}

}
