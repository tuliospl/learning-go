package main

import "fmt"

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

// Recursive functions call themselves until a base case is reached.
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// The recover function can be used to handle panics and prevent the program from crashing.
func recoverExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("This will be printed first")
	panic("Something went wrong!")
	fmt.Println("This will not be printed because of the panic")
}

// Calculator Methods are functions that are associated with a type. They can be defined on structs and can have a receiver.
type Calculator struct {
	value int
}

func (c *Calculator) Add(n int) {
	c.value += n
}

func (c Calculator) GetValue() int {
	return c.value
}

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
	fmt.Println()

	fmt.Printf("Factorial of 5: %d\n", factorial(5))
	fmt.Println()

	recoverExample()
	fmt.Println()

	calc := &Calculator{value: 10}
	calc.Add(5)
	fmt.Printf("Calculator value: %d\n", calc.GetValue())
}
