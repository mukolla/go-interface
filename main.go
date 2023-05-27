package main

import "fmt"

type animal interface {
	walker
	runner
}

type birg interface {
	walker
	flier
}

type walker interface {
	walk()
}

type runner interface {
	run()
}

type flier interface {
	fly()
}

type cat struct {
}

type eagle struct {
}

type stork struct {
}

func (c *cat) walk() {
	fmt.Println("Cat is walking")
}

func (c *cat) run() {
	fmt.Println("Cat is running")
}

func (e *eagle) walk() {
	fmt.Println("Eagle is walk")
}

func (e *eagle) fly() {
	fmt.Println("Eagle is fly")
}

func (e *stork) walk() {
	fmt.Println("Eagle is walk")
}

func (e *stork) fly() {
	fmt.Println("Eagle is fly")
}

func walk(w walker) {
	w.walk()
}

func main() {
	var c animal = &cat{}
	var e birg = &eagle{}
	var s birg = &stork{}

	fmt.Print("\n")
	walk(c)
	walk(e)
	walk(s)
	fmt.Print("\n")
}
