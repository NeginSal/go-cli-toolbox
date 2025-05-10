package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func unitConverter() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Unit Converter")

	fmt.Println("Choose the desired unit:")
	fmt.Println("1. Meter to Kilometer")
	fmt.Println("2. Kilogram to Gram")
	fmt.Println("3. Celsuis to Fahrenheit")

	fmt.Print("Your choice (number):")

	unitInput, _ := reader.ReadString('\n')
	unitInput = strings.TrimSpace(unitInput)

	fmt.Print("Enter the value:")
	valueInput, _ := reader.ReadString('\n')
	valueInput = strings.TrimSpace(valueInput)
	value, err := strconv.ParseFloat(valueInput, 64)

	if err != nil {
		fmt.Println("❌ Invalid value.")
		return
	}

	switch unitInput {
	case "1":
		result := value / 1000
		fmt.Printf("%.2f meters is equal to %.4f  kilometers\n", value, result)
	case "2":
		result := value * 1000
		fmt.Printf("%2f kilograms is equal to %.0f grams \n", value, result)
	case "3":
		result := (value * 9 / 5) + 32
		fmt.Printf("%.2f degrees Celsius is equal to %.2f degrees Fahrenheit\n", value, result)
	default:
		fmt.Println("❌ Invalid option.")
	}
}
