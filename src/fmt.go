package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	fmt.Println("Hello World!") //  Hello World!

	fmt.Println()

	fmt.Printf("%t\n", true) //  true
	fmt.Printf("%d\n", 64)   //  64
	fmt.Printf("%b\n", 18)   //  10010
	fmt.Printf("%c\n", 78)   //  N
	fmt.Printf("%x\n", 12)   //  c

	fmt.Println()

	point := point{1, 2}
	fmt.Printf("%v\n", point)  // {1 2}
	fmt.Printf("%+v\n", point) // {x:1 y:2}
	fmt.Printf("%#v\n", point) // main.point{x:1, y:2}    Note: Type available at runtime
	fmt.Printf("%T\n", point)  // main.point

	fmt.Println()

	fmt.Printf("%f\n", 7445.23)    // 7445.230000
	fmt.Printf("%e\n", 98760000.0) // 9.876000e+07

	fmt.Println()

	fmt.Printf("%s\n", "\"string\"") // "string"
	fmt.Printf("%q\n", "\"string\"") // "\"string\""
	fmt.Printf("%x\n", "cherry lab") // 636865727279206c6162

	fmt.Println()

	fmt.Printf("%p\n", &point) // Non-determinant stack location

	fmt.Printf("|%6d|%6d|\n", 12, 345)         // |    12|   345|
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)   // |  1.20|  3.45|
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45) // |1.20  |3.45  |
}
