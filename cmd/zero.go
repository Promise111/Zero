package cmd

import (
	"fmt"
	"strconv"
)

func Add(first, second string) (result string) {
	firstFloat, err := strconv.ParseFloat(first, 64)
	if err != nil {
		fmt.Println("Error: First value is invalid")
		return
	}

	secondFloat, err := strconv.ParseFloat(second, 64)
	if err != nil {
		fmt.Println("Error: Second value is invalid")
		return
	}

	return fmt.Sprintf("%f", firstFloat+secondFloat)
}

func Subtract(first, second string) (result string) {
	firstFloat, err := strconv.ParseFloat(first, 64)
	if err != nil {
		fmt.Println("Error: First value is invalid")
		return
	}

	secondFloat, err := strconv.ParseFloat(second, 64)
	if err != nil {
		fmt.Println("Error: Second value is invalid")
		return
	}

	return fmt.Sprintf("%f", firstFloat-secondFloat)
}

func Multiply(first, second string, shouldRoundUp bool) (result string) {
	firstNum, err := strconv.ParseFloat(first, 64)
	if err != nil {
		fmt.Println("Error: First value is not a decimal")
		return
	}

	secondNum, err := strconv.ParseFloat(second, 64)
	if err != nil {
		fmt.Println("Error: Second value is not a decimal")
		return
	}

	if shouldRoundUp {
		return fmt.Sprintf("%.2f", firstNum*secondNum)
	}

	return fmt.Sprintf("%f", firstNum*secondNum)
}

func Divide(first, second string, shouldRoundUp bool) (result string) {
	firstNum, err := strconv.ParseFloat(first, 64)
	if err != nil {
		fmt.Println("Error: First value is invalid")
		return
	}

	secondNum, err := strconv.ParseFloat(second, 64)
	if err != nil {
		fmt.Println("Error: Second value is invaid")
		return
	}

	if secondNum == 0 {
		fmt.Println("Error: Division by zero")
		return
	}

	if shouldRoundUp {
		return fmt.Sprintf("%.2f", firstNum/secondNum)
	}

	return fmt.Sprintf("%f", firstNum/secondNum)
}
