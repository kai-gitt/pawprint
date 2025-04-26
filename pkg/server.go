package pkg

import "fmt"

type serverT struct{}

var Server serverT

func (serverT) Hello() {
	fmt.Println("hello, world")
}
