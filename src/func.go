package main

import (
	"fmt"
	"sort"
)

func sum(a, b int) int { // Return type put after the parameters
	return a + b
}

func meanAndMedian(n ...float64) (float64, float64) { // Ellipsis to define listed inputs. Function returns tuple
	total := 0.0
	for _, v := range n { // range keyword iterates through will tuple that is (index, value)
		total += v
	}
	count := len(n)
	mean := total / float64(count)
	sort.Float64s(n)
	middleIdx := count / 2 // Integer division takes the floor
	var median float64
	if middleIdx%2 == 0 {
		median = (n[middleIdx-1] + n[middleIdx]) / 2.0
	} else {
		median = n[middleIdx]
	}
	return mean, median
}

func main() {
	i := sum(23, 20)
	fmt.Println(i)

	mean, median := meanAndMedian(1.0, 2.0, 6.0, 8.0) // Unpacking similar to python
	fmt.Println(mean, median)
}
