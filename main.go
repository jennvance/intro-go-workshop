package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type record struct {
	HighScore   int
	Data        map[string]int
	ListOfScore []int
}

func (r *Record) Calculate() int {
	return 42
}

func main() {
	r := &record{HighScore: 1}
	rand.Seed(time.Now().Unix())
	fmt.Println("Welcome to guessing game")
	//declare and assign variable at same time
	secretNumber := rand.Intn(10)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Guess number between 0 and 9")
		//ask for input
		scanner.Scan()

		input, err := scanner.Text(), scanner.Err()
		if err != nil {
			fmt.Println("What?")
			continue
		}

		//convert string to int
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Not a number.")
			continue
		}

		if num == secretNumber {
			fmt.Println("YAY. Win!")
			break
		}
		fmt.Println("Nope. Try again.")
	}

	fmt.Println("GAME OVER")
}
