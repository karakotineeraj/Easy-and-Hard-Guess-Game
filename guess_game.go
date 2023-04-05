package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fmt.Println("GUESS GAME: ")
	fmt.Print("\t Press 1 for Easy and 2 for Hard: ")

	var mode int
	fmt.Scan(&mode)

	if mode == 1 {
		easyMode()
	} else {
		hardMode()
	}
}

func hardMode() {
	randomNumber := rand.Intn(1001)
	var guess = -1
	tries := 1

	for guess != randomNumber {
		fmt.Print("Enter your guess (between 0 - 1000 included): ")
		fmt.Scan(&guess)

		if guess > randomNumber {
			fmt.Println("Guess Lower")
		} else if guess < randomNumber {
			fmt.Println("Guess Higher")
		} else {
			fmt.Printf("Congrats, the number was %d. You guessed it in %d tries.", randomNumber, tries)
			break
		}

		tries++
	}
}

func easyMode() {
	randomNumber := rand.Intn(1001)
	var guess = -1
	tries := 1

	low, high := 0, 1000

	for guess != randomNumber {
		fmt.Printf("Enter your guess (between %d - %d included): ", low, high)
		fmt.Scan(&guess)

		if guess > randomNumber {
			high = int(math.Min(float64(guess-1), float64(high)))
		} else if guess < randomNumber {
			low = int(math.Max(float64(guess+1), float64(low)))
		} else {
			fmt.Printf("Congrats, the number was %d. You guessed it in %d tries.", randomNumber, tries)
			break
		}

		tries++
	}
}
