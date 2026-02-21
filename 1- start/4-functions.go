package main

import "fmt"

func main() {
	checkIsZero(0)
	checkIsZero(1)
	fmt.Println()

	subtraction, multiplication := calculate(5, 10)
	fmt.Printf("Subtraction: %d, Multiplication: %d\n", subtraction, multiplication)
	fmt.Println()

	result1, err := division(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Division Result: %d\n", result1)
	}
	result2, err := division(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Division Result: %d\n", result2)
	}
	fmt.Println()

	total := sum(1, 2, 3, 4, 5)
	fmt.Printf("Total Sum: %d\n", total)
	fmt.Println()

	greeting()
	fmt.Println()

	multiply := func(x, y int) int {
		return x * y
	}
	fmt.Printf("Total Multiply: %d\n\n", applyOperation(5, 3, multiply))
	fmt.Println()

	decrement := countdown(5)
	fmt.Println("Countdown:", decrement())
	fmt.Println("Countdown:", decrement())
	fmt.Println("Countdown:", decrement())
	fmt.Println("Countdown:", decrement())
	fmt.Println("Countdown:", decrement())
	fmt.Println()

	deferExample()
}

// Functions can return a value
func checkIsZero(param int) {
	fmt.Println(param > 0)
	return
}

// Functions can return multiple values
func calculate(a int, b int) (subtraction int, multiplication int) {
	subtraction = a - b
	multiplication = a * b
	return
}

// Functions can also return an error
func division(a int, b int) (result int, err error) {
	if b == 0 {
		err = fmt.Errorf("cannot divide by zero")
		return
	}
	result = a / b
	return
}

// Variadic functions can take a variable number of arguments
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// Functions can be defined inside other functions
func greeting() {
	greeting := func(name string) string {
		return "Hello, " + name
	}
	fmt.Println(greeting("Tulio"))
}

// Functions can be passed as arguments to other functions (higher-order functions)
func applyOperation(a int, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// Functions can return other functions (closures)
func countdown(n int) func() int {
	counter := n
	return func() int {
		if counter > 0 {
			counter--
			return counter
		}
		fmt.Println("Blast off!")
		return 0
	}
}

// Deferred functions are executed in LIFO order.
func deferExample() {
	defer fmt.Println("This will be printed last because it's deferred")
	fmt.Println("This will be printed first")
}
