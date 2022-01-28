package main

import "fmt"

type mo float64

var x mo
var y int

func main() {
	x = 45.6
	fmt.Printf("\n%v\n", x)
	y = int(x)
	fmt.Printf("\n%v\n", y)
}
