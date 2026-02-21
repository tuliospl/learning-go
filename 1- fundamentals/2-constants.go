package main

// Constants are immutable values that cannot be changed after they are defined
const pi = 3.1415
const gravity = 9.81
const learningGo = "Go is fun!"
const isGoAwesome = true

// You can also define multiple constants in a block
const (
	a = 1
	b = 2
	c = "tres"
)

func main() {
	println("Value of pi:", pi)
	println("Value of gravity:", gravity)
	println("Value of learningGo:", learningGo)
	println("Value of isGoAwesome:", isGoAwesome)
	println("Value of a:", a)
	println("Value of b:", b)
	println("Value of c:", c)
}
