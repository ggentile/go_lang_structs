package main

import "fmt"

func main() {
	stack := Stack{}

	stack.Push("Tiago")
	stack.Push("Daniel")
	stack.Push("Giovanna")

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

}
