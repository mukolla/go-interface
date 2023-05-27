package main

import "fmt"

type bird interface {
	walker
	flier
}

type walker interface {
	walk()
}

type flier interface {
	fly()
}

type flyImpl struct {
	name string
}

func (f *flyImpl) fly() {
	fmt.Printf("%s is flying\n", f.name)
}

type eagle struct {
	walker
	flyImpl
}

type stork struct {
	walker
	flyImpl
}

func (e *eagle) walk() {
	fmt.Printf("%s is walking\n", e.name)
}

func (s *stork) walk() {
	fmt.Printf("%s is walking\n", s.name)
}

func main() {
	var b bird

	e := &eagle{
		flyImpl: flyImpl{name: "Eagle"},
	}
	s := &stork{
		flyImpl: flyImpl{name: "Stork"},
	}

	b = e
	b.walk()
	b.fly()

	b = s
	b.walk()
	b.fly()
}
