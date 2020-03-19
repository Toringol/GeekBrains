package main

import "fmt"

type AgingInterface interface {
	AgeUp()
}

type Human struct {
	Name string
	Age  int
}

func (h *Human) AgeUp() {
	h.Age++
}

type Animal struct {
	Name string
	Age  int
}

func (an *Animal) AgeUp() {
	an.Age++
}

func main() {
	h := &Human{
		Name: "Sergey",
		Age:  22,
	}
	an := &Animal{
		Name: "Bird",
		Age:  6,
	}
	AgingInterface.AgeUp(h)
	AgingInterface.AgeUp(an)
	fmt.Println(h)
	fmt.Println(an)
}
