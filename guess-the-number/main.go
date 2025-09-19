package main

import (
	"fmt"
	"math/rand"
)

func generateSecretNumber() int {
	fmt.Println("Generating secret number, the secret number higher than 0 and lower than 10")
	secretNumber := rand.Intn(10)
	fmt.Println("Secret Number was generated")
	return secretNumber
}

func enterNumber() int {
	fmt.Println("Enter number")
	var number int
	_, err := fmt.Scanln(&number)
	if err != nil {
		return -1
	}
	return number
}

func validateNumber(number int, secretNumber int) {
	if number < secretNumber {
		fmt.Println("The secret number is lower than the secret number")
	} else if number > secretNumber {
		fmt.Println("The secret number is higher than the secret number")
	} else {
		fmt.Println("Wow you guessed correctly")
	}
}

func run() {
	secretNumber := generateSecretNumber()
	number := -1
	for number != secretNumber {
		number = enterNumber()
		validateNumber(number, secretNumber)
	}
}

func main() {
	fmt.Println("Welcome to Guess The Number Game")
	run()
}
