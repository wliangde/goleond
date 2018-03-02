package test

import "fmt"

type Tester interface {
	Main()
}

func Test1() {
	fmt.Println("hello  world!")
}
