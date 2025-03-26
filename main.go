package main

import "fmt"

func main() {
	var secondNumber float64
	var firstNumber float64
	var operation string
	var result float64
	fmt.Println("Hi. This is calculator\n Choose your First Number to take math operation")
	fmt.Scan(&firstNumber)
	fmt.Println("What is your Second Number?")
	fmt.Scan(&secondNumber)
	fmt.Println("Choose operation?")
	fmt.Scan(&operation)
	switch operation {
	case "+":
		result = firstNumber + secondNumber
	case "-":
		result = firstNumber - secondNumber
	case "*":
		result = firstNumber * secondNumber
	case "/":
		if secondNumber == 0 {
			fmt.Println("You can't divide to zero!")
			break
		}
		result = firstNumber / secondNumber
	}
	fmt.Println(result)
}
