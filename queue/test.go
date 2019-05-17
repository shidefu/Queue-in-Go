package main

import "fmt"

func main() {
	var queue Queue
	queue.Offer(1)
	queue.Offer(2)
	queue.Offer(3)
	a, _ := queue.Poll()
	b, _ := queue.First()
	queue.Offer(4)
	c, _ := queue.Last()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(queue.Size())
}
