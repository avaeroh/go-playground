package common

import (
	"fmt"
	"math/rand"
	"time"
)

func generateNum() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 3
	return rand.Intn(max-min+1) + min
}

func countdown(duration int) {
	for duration > 0 {
		fmt.Println(duration, "...")
		time.Sleep(1 * time.Second)
		duration--
	}
}

//very lazy, should do this in a smarter way
func CompareSelection(userChoice int) bool {
	playerWin := false
	compChoice := generateNum()
	if userChoice == generateNum() {
		countdown(3)
		fmt.Println("DRAW!")
		return playerWin
		//rock == 1, paper == 2, scis == 3
	} else if compChoice == userChoice+1 {
		playerWin = false
	} else if userChoice == compChoice+1 {
		playerWin = true
	} else if userChoice < compChoice {
		playerWin = false
	} else {
		playerWin = true
	}
	fmt.Printf("You chose %v, the computer chose...\n", numToChoice(userChoice))
	countdown(3)
	fmt.Printf(numToChoice(userChoice) + "!")
	return playerWin
}

func numToChoice(choice int) string {
	if choice == 1 {
		return "rock"
	} else if choice == 2 {
		return "paper"
	} else {
		return "scissors"
	}
}
