package main

import (
	"crud_product/cmd"
	"fmt"
)

func main() {
	cmd.RunServe()
}

type A struct {
	Name string
	Age  int
}

type B struct {
	Name string
	Age  int
}

func testing() {
	a := A{
		Name: "User 1",
		Age:  19,
	}
	var b B

	b = B{
		Name: a.Name,
		Age:  a.Age,
	}

	fmt.Println(a)
	fmt.Println(b)
}
