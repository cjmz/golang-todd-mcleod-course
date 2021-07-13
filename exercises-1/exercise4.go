package main

import "fmt"

type jeremias int

var x jeremias

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

}
