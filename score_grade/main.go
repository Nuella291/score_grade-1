package main

import (
	"fmt"
	"score_grade/grader"
	"strings"
)

func main() {
	for {
		var studentName string
		var studentScore float64
		var grade string
		var option string

		fmt.Print("Enter your name: \n")
		fmt.Scan(&studentName)

		fmt.Printf("Hello, %v! Please enter your score: \n", strings.Title(studentName))
		fmt.Scan(&studentScore)

		grade = grader.DisplayGrade(studentScore)
		fmt.Println("Your grade is: ", grade)

		fmt.Print("Do you want to continue? (yes/no): ")
		fmt.Scan(&option)

		option = strings.ToLower(strings.TrimSpace(option))

		if option == "no" || option == "n" {
			fmt.Println("Program ended.")
			break
		}

		fmt.Println()
	}
}
