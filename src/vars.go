package main

import "fmt"

func main() {
	a := 5 // Type inference
	fmt.Println(a)

	var str = "My String"
	fmt.Println(str)

	var b int
	b = 2
	fmt.Println(b)

	var c, d = 3, 4
	fmt.Printf("%d, %d\n", c, d)

	var e int
	fmt.Println(e)

	slice := []int { // Variable size 'slice' placed on heap. Size can be modified
		1, 2, 3,
	}
	fmt.Println(slice)

	arr := [3]int { // Fixed size array placed on stack. Block of memory with size for three integers only
		4, 5, 6,
	}
	fmt.Println(arr)

	myMap := map[string]int { // map[Key]Value
		"BadRequest": 120,
		"NoAuth": 65,
		"OK": 30,
	}
	fmt.Println(myMap)
}
