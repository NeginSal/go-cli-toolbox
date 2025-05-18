package main

import "github.com/fatih/color"

var (
	Red     = color.New(color.FgRed).SprintFunc()
	Green   = color.New(color.FgGreen).SprintFunc()
	Yellow  = color.New(color.FgYellow).SprintFunc()
	Cyan    = color.New(color.FgCyan).SprintFunc()
	Blue    = color.New(color.FgBlue).SprintFunc()
	White   = color.New(color.FgWhite).SprintFunc()
	Magenta = color.New(color.FgMagenta).SprintFunc()
)

func Error(msg string) string {
	return Red("❌ " + msg)
}

func Success(msg string) string {
	return Green("✅ " + msg)
}

func Warning(msg string) string {
	return Yellow("⚠️ " + msg)
}

func Info(msg string) string {
	return Cyan("ℹ️ " + msg)
}

func SimpleTextOne(msg string) string {
	return Blue(msg)
}

func SimpleTextTwo(msg string) string {
	return Magenta(msg)
}
