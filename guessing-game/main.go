package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(100)
	fmt.Println("Secret Number is ", secretNumber)

	fmt.Println("Guess a number between 1 and 100")
	fmt.Println("Enter your guess:")

	attempts := 0
	for {
		attempts++
		reader := bufio.NewReader(os.Stdin)

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error while reading the input -", err)
			continue
		}

		input = strings.TrimSuffix(input, "\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Entered wrong input, need an integer value -", err)
			continue
		}

		fmt.Println("Your guess is : ", guess)

		if guess > secretNumber {
			fmt.Println("Guess is higher than secret number. Enter another guess")
		} else if guess < secretNumber {
			fmt.Println("Guess is smaller than secret number. Enter another guess")
		} else {
			fmt.Println("Congrts your guess has matched scret number !! in ", attempts, "attempts")
			break
		}
	}
}
