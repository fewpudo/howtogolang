package main

import "fmt"

type mo int

var x mo

func main() {
	x = 20
	fmt.Printf("%d", x)
	fmt.Printf("\n %T ", x)
	x = 42
	fmt.Printf("\n %v", x)
}
