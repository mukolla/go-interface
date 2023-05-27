package main

import "fmt"

type greeter interface {
	greet(string) string
}

type ukrainian struct {
}

type english struct {
}

func (u *ukrainian) greet(name string) string {
	return fmt.Sprintf("Привіт, %s", name)
}

func (e *english) greet(name string) string {
	return fmt.Sprintf("Привіт, %s", name)
}

func sayHello(g greeter, name string) {
	fmt.Println(g.greet(name))
}

func main() {
	fmt.Print("\n")
	sayHello(&ukrainian{}, "Пилип")
	sayHello(&english{}, "Bill")
	fmt.Print("\n")
}
