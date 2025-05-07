package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bmi() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("BMI Calculation (Body Mass Index)")
	fmt.Println("Enter your weight in kilogram:")

	weightStr, _ := reader.ReadString('\n')
	weight, err := strconv.ParseFloat(strings.TrimSpace(weightStr), 64)

	if err != nil || weight <= 0 {
		fmt.Println("❌ Invalid weight.")
		return
	}

	fmt.Print("Enter your height in centimeters: ")
	heightStr, _ := reader.ReadString('\n')
	height, err := strconv.ParseFloat(strings.TrimSpace(heightStr), 64)

	if err != nil || height <= 0 {
		fmt.Println("❌ Invalid height.")
		return
	}

	heightM := height / 100
	bmi := weight / (heightM * heightM)

	fmt.Printf("✅ Your BMI: %.2f\n", bmi)

	switch {
	case bmi < 18.5:
		fmt.Println("Status: Underweight")
	case bmi >= 18.5 && bmi < 25:
		fmt.Println("Status: Normal")
	case bmi >= 25 && bmi < 30:
		fmt.Println("Status: Overweight")
	default:
		fmt.Println("Status: Obese")
	}
}
