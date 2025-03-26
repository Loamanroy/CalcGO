package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	for again := true; again; {
		var firstNumber, secondNumber, result float64
		var operation, continueInput string
		fmt.Println("Hi. This is calculator\nChoose your First Number to take math operation")
		if _, err := fmt.Scan(&firstNumber); err != nil {
			fmt.Println("Error reading input:", err.Error())
			bufio.NewReader(os.Stdin).ReadString('\n')
			continue
		}
		fmt.Println("What is your Second Number?")
		if _, err := fmt.Scan(&secondNumber); err != nil {
			fmt.Println("Error reading input:", err.Error())
			bufio.NewReader(os.Stdin).ReadString('\n')
			continue
		}
		fmt.Print("Choose operation (+, -, *, /, ^, %, sqrt): ")
		if _, err := fmt.Scan(&operation); err != nil {
			fmt.Println("Error reading operation:", err.Error())
			bufio.NewReader(os.Stdin).ReadString('\n')
			continue
		}
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
				continue
			}
			result = firstNumber / secondNumber
		case "^":
			result = math.Pow(firstNumber, secondNumber)
		case "%":
			result = float64(int32(firstNumber) % int32(secondNumber))
		case "sqrt":
			fmt.Println("You are chose sqrt operation, so your secondNumber was reduced to 0")
			result = math.Sqrt(firstNumber)
			secondNumber = 0
		default:
			fmt.Println("Invalid operation")
			continue
		}
		fmt.Println("Result is:", result)
		fmt.Print("Again (y/n)?:")
		fmt.Scan(&continueInput)
		switch continueInput {
		case "y":
			again = true
		case "n":
			again = false
		default:
			fmt.Println("Incorrect value")
			again = false
		}
	}

}
