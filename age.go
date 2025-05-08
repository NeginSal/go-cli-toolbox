package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func age() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Age Calculation")

	fmt.Print("Enter your birth date (Format:YYYY-MM-DD): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Parse string to date format
	birthDate, err := time.Parse("2006-01-02", input)
	if err != nil {
		fmt.Println("❌ Invalid date format. Please enter in YYYY-MM-DD format.")
		return
	}

	today := time.Now()

	// Calculate the time difference
	years := today.Year() - birthDate.Year()
	months := int(today.Month()) - int(birthDate.Month())
	days := today.Day() - birthDate.Day()

	// Correct values if they are negative
	if days < 0 {
		months--
		days += 30 // Approximately, because the number of days in months varies
	}
	if months < 0 {
		years--
		months += 12
	}

	fmt.Printf("✅ You are %d years, %d months, and %d days old.\n", years, months, days)
}