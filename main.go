package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {
		fmt.Println("Available tools")
		fmt.Println("1. Simple Calculator")
		fmt.Println("2. Calculate BMI")
		fmt.Println("3. Time conversion (local time <-> UTC)")
		fmt.Println("4. Calculate age")
		fmt.Println("5. Unit conversion (e.g. , kg <-> lb)")
		fmt.Println("0. Exit")

		fmt.Println("Please enter your option: ")

		// reading input from the user
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		choice, err := strconv.Atoi(strings.TrimSpace(input))

		if err != nil {
			fmt.Println("❌ Invalid input. Please enter only numbers.")
			continue
		}

		// checking the entered option
		switch choice {
		case 1:
			calculator()
		case 2:
			bmi()
		case 3:
			fmt.Println("Time conversion selected.")
		case 4:
			fmt.Println("Age calculation selected.")
		case 5:
			fmt.Println("Unit conversion selected.")
		case 0:
			fmt.Println("Exiting the program. Goodbye! 👋")
			return
		default:
			fmt.Println("❌ Invalid option. Please enter a valid number.")
		}
	}
}
