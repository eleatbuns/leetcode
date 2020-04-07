package main

import "fmt"

type IFoo interface {
	Write() (n int, err error)
	Read() (n int, err error)
}

type Foo struct {
	IFoo
}
type Foc struct { // Go 文法
	IFoo
}

var foo IFoo = new(Foo)
var foc IFoo = new(Foc)

func main() {
	foo.Write()
}


func (Foo) Write() (n int, err error) {
	fmt.Println("FOO.Write")
	return 0, nil
}

func (Foc) Write() (n int, err error) {
	fmt.Println("FOC.Write")
	return 0, nil
}
