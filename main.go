package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func mult(a int, b int) int {
	return a * b
}

func div(a int, b int) int {
	return a / b
}

func main() {
	fmt.Println("Calculadora #GoEssentials")

	result := sum(5, 10)
	fmt.Println(result)
	result = sub(50, 10)
	fmt.Println(result)
	result = div(10, 5)
	fmt.Println(result)
	result = mult(5, 100)
	fmt.Println(result)

}
