package main

import "fmt"

func main() {
	fmt.Println(variableNoDeclarada)

	var x int = "esto es un string"

	resultado := suma("5", 3)
	fmt.Println(resultado)
}
func suma(a string, b int) string {
	return a + b
}
