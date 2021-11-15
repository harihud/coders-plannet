package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	foo()
	for i := 0; i < 10; i++ {
		print(i)
	}
	bar()
	print("Exiting...!")
}

func foo() {
	print("I'm in foo.")
}

func bar() {
	print("I'm in bar")
}
