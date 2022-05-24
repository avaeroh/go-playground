package main

import (
	"common"
	"fmt"
)

var user string
var userChoice int
var replay bool = true
var compScore int
var userScore int

func main() {

	//setup
	fmt.Println("Let's play a CLI game that I'll decide what it does as I write it")
	fmt.Println("Firstly, what is your name?")
	for user == "" {
		fmt.Scanln(&user)
		if user == "" {
			fmt.Println("Try that again.")
		}
	}
	fmt.Printf("Hello %v!\n", user)

	for replay {

		//choice
		fmt.Println("Rock (1), paper (2), scissors (3)?")
		for {
			_, err := fmt.Scanln(&userChoice)
			if (userChoice > 3 || userChoice < 0) || err != nil {
				fmt.Println("Try that again.")
			} else {
				break
			}
		}

		//drumroll
		playerWin := common.CompareSelection(userChoice)
		if playerWin {
			fmt.Printf("Congrats %v, you won!\n", user)
			userScore++
			fmt.Println("Current score:")

		} else {
			fmt.Printf("Unlucky %v, you lost!\n", user)
			compScore++
		}
		fmt.Printf("%v: %d, Computer: %d\n", user, userScore, compScore)

		if userScore > compScore {
			fmt.Println("One more? You can't finish on a win...")
		} else if userScore < compScore {
			fmt.Println("One more? You can't finish on a loss...")
		} else {
			fmt.Println("One more? You can't finish on a tie...")
		}

		for replay == true {
			var input string
			fmt.Println("Replay (y) or quit (q)?")
			fmt.Scanln(&input)
			if input != "y" && input != "n" {
				fmt.Println("Try that again.")
			}
			if input == "q" {
				replay = false
				break
			}
			if input == "y" {
				fmt.Println("Alright, here we go again...")
				break
			}
		}

	}

}
