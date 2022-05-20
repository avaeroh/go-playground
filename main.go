package main

import (
	adele "common"
	t "fmt"
)

func main() {
	t.Println("hey!")
	a := "I don't need to tell anyone I'm a"
	t.Printf("%v %T \n", a, a)
	adele.HelloFromTheOtherSide()
}
