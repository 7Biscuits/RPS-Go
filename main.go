package main

import (
	"fmt"
	"strings"
	"math/rand"
	"time"
)

var gameElements [3] string
var userMove string
var userPoints int = 0
var compPoints int = 0

func main() {
	gameElements[0] = "rock"
	gameElements[1] = "paper"
	gameElements[2] = "scissor"
	
	var gameRounds int
	fmt.Println("Enter the number of rounds: ")
	fmt.Scan(&gameRounds)
	for i := 0; i < gameRounds; i++ {
		gameLogic()
	}
	fmt.Println("User's score : ", userPoints)
	fmt.Println("Computer's score : ", compPoints)
	if userPoints > compPoints {
		fmt.Println("User wins")
	} else if compPoints > userPoints {
		fmt.Println("Computer wins")
	} else {
		fmt.Println("It's a draw")
	}
}

func gameLogic() {
	var result string
	fmt.Println("Your move : ")
	fmt.Scan(&userMove)
	rand.Seed(time.Now().UnixNano())
	compMove := randInt(0, 3)
	if strings.ToLower(userMove) == gameElements[compMove] {
		result = "draw"
	} else if strings.ToLower(userMove) == "rock" && gameElements[compMove] == "paper" || strings.ToLower(userMove) == "paper" && gameElements[compMove] == "scissor" || strings.ToLower(userMove) == "scissor" && gameElements[compMove] == "rock"{
		result = "computer wins"
		compPoints++
	} else if strings.ToLower(userMove) == "rock" && gameElements[compMove] == "scissor" || strings.ToLower(userMove) == "paper" && gameElements[compMove] == "rock" || strings.ToLower(userMove) == "scissor" && gameElements[compMove] == "paper"{
		result = "user wins"
		userPoints++
	}  else {
		fmt.Println("Invalid Input")
		return 
	}
	fmt.Println("User Move: ", strings.ToLower(userMove))
	fmt.Println("Computer Move: ", gameElements[compMove])
	fmt.Println(result)
}

func randInt(min, max int) int {
    return min + rand.Intn(max-min)
}