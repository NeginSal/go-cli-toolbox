package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	persian "github.com/yaa110/go-persian-calendar"
)

// Convert  ShamsiToGregorian
func gegorianrToshamsi(t time.Time) string {
	p := persian.New(t)
	return fmt.Sprintf("%04d/%02d/%02d", p.Year(), p.Month(), p.Day())
}

// Convert Shamsi to Gregorian
func shamsiToGregorian(year, month, day int) (time.Time, error) {
	p := persian.New(time.Now())                                 // Create a new Persian calendar object
	p.Set(year, persian.Month(month), day, 0, 0, 0, 0, time.UTC) // Set the Shamsi date with all components
	return p.Time(), nil                                         // Get the ptime.Time and convert it to Go's time.Time
}

func dateConverter() {
	var choice int
	cyan := color.New(color.FgCyan).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()

	fmt.Println(cyan("📅 Select conversion type:"))
	fmt.Println()
	fmt.Println(blue("1: Gregorian → Shamsi"))
	fmt.Println(blue("2: Shamsi → Gregorian"))
	fmt.Println()
	fmt.Print(magenta("*Your choice: "))
	fmt.Scanln(&choice)
	fmt.Println()

	switch choice {
	case 1:
		var year, month, day int
		fmt.Print(yellow("📆 Gregorian date (year month day): "))
		fmt.Scanln(&year, &month, &day)
		t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
		fmt.Println()
		fmt.Println(green("🟢 Shamsi:"), gegorianrToshamsi(t)) // Print the Shamsi date
		fmt.Println()

	case 2:
		var year, month, day int
		fmt.Print(yellow("📆 Shamsi date (year month day): "))
		fmt.Scanln(&year, &month, &day)
		t, err := shamsiToGregorian(year, month, day)
		if err != nil {
			fmt.Println(red("❌ Conversion error:"), err)
			return
		}
		fmt.Println()
		fmt.Println(green("🟢 Gregorian:"), t.Format("2006-01-02")) // Print the Gregorian date
		fmt.Println()

	default:
		fmt.Println(red("❌ Invalid choice."))
	}
}
