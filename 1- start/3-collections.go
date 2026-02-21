package main

import "fmt"

// Arrays fixed size
var array1 [3]int
var array2 = [5]int{1, 2, 3, 4, 5}

// Slices dynamic size
var slice1 []int
var slice2 = []int{10, 20, 30, 40, 50}

// Maps key-value pairs
var map1 map[string]int
var map2 = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
}

// Structs custom data types
var struct1 struct{}
var struct2 = struct {
	name string
	age  int
}{
	name: "Alice",
	age:  25,
}

func main() {
	fmt.Println("Array:", array1)
	fmt.Println("Array:", array2)
	fmt.Println()

	fmt.Println("Slice:", slice1)
	fmt.Println("Slice:", slice2)
	fmt.Println()

	fmt.Println("Map:", map1)
	fmt.Println("Map:", map2)
	fmt.Println()

	fmt.Println("Struct:", struct1)
	fmt.Println("Struct:", struct2)
}
