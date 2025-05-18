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
		fmt.Println(Magenta(" 🛠️  Available tools:"))
		fmt.Println()
		fmt.Println(Cyan(" 1. Simple Calculator"))
		fmt.Println(Cyan(" 2. Calculate BMI "))
		fmt.Println(Cyan(" 3. Time conversion (local time <-> UTC) "))
		fmt.Println(Cyan(" 4. Calculate age "))
		fmt.Println(Cyan(" 5. Unit conversion (e.g. , kg <-> lb) "))
		fmt.Println(Cyan(" 0. Exit "))
		fmt.Println()
		fmt.Println(Magenta("Please enter your option: "))

		// reading input from the user
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		choice, err := strconv.Atoi(strings.TrimSpace(input))

		if err != nil {
			Red("❌ Invalid input. Please enter only numbers.")
			continue
		}

		// checking the entered option
		switch choice {
		case 1:
			calculator()
		case 2:
			bmi()
		case 3:
			dateConverter()
		case 4:
			age()
		case 5:
			unitConverter()
		case 0:
			fmt.Println()
			Green("Exiting the program. Goodbye! 👋")
			return
		default:
			Red("❌ Invalid option. Please enter a valid number.")
		}
	}
}
