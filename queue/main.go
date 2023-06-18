package main

import "fmt"

func main() {

	queue := Queue{}

	queue.Enqueue("Tiago")
	queue.Enqueue("Maria")
	queue.Enqueue("João")
	queue.Enqueue("José")

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())

}
