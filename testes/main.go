package main

import "fmt"

type yuan string

var z yuan

func main() {
	z = "Vai pro caralho"
	fmt.Printf("%T", z)
	fmt.Println(z)
}
