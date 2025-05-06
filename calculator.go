package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculator() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Simple Calculator")
	fmt.Print("Enter your first number: ")
	num1Str, _ := reader.ReadString('\n')
	num1, err := strconv.ParseFloat(strings.TrimSpace(num1Str), 64)

	if err != nil {
		fmt.Println("❌ The first number is not valid.")
		return
	}

	fmt.Print("Enter your second number: ")
	num2Str, _ := reader.ReadString('\n')
	num2, err := strconv.ParseFloat(strings.TrimSpace(num2Str), 64)

	if err != nil {
		fmt.Println("❌ The second number is not valid.")
		return
	}

	fmt.Print("Enter your operation (+, -, *, /):")
	op,_ := reader.ReadString('\n')
	op = strings.TrimSpace(op)

	var result float64
	valid :=true

	switch op {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
  case "*":
	result = num1 * num2
	case "/":
		if num2 ==0 {
			fmt.Println("❌ Division by zero is not allowed.")
			return
		}
		result = num1 /num2
	default:
		valid = false
	}

	if valid {
		fmt.Printf("✅ Result: %.2f\n", result) 
	} else {
		fmt.Println("❌ The entered operation is not valid.") 
	}
}
