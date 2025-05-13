package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func main() {
	for {
		color.Cyan(" 🛠️  Available tools:")
		fmt.Println()
		color.Yellow(" 1. Simple Calculator")
		color.Yellow(" 2. Calculate BMI ")
		color.Yellow(" 3. Time conversion (local time <-> UTC) ")
		color.Yellow(" 4. Calculate age ")
		color.Yellow(" 5. Unit conversion (e.g. , kg <-> lb) ")
		color.Yellow(" 0. Exit ")
		fmt.Println()
		color.Cyan("Please enter your option: ")

		// reading input from the user
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		choice, err := strconv.Atoi(strings.TrimSpace(input))

		if err != nil {
			color.Red("❌ Invalid input. Please enter only numbers.")
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
			color.Green("Exiting the program. Goodbye! 👋")
			return
		default:
			color.Red("❌ Invalid option. Please enter a valid number.")
		}
	}
}
