package main

import "fmt"

// import "practice"
import "learn/stack"

func main() {
	s := new(stack.Stack)

	s.Push(1)
	s.Push(3)
	s.Push(2)
	fmt.Printf("%v\n", s)

	s.Pop()
	fmt.Printf("%v\n", s)

}
