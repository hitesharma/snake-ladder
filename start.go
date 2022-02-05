package main

import "fmt"

func StartGame() {
	for {
		currentPlayer, err := Dequeue()
		if err != nil {
			fmt.Println("Game Over")
			break
		}
		diceValue := RollDice()
		currentPlayer.position += diceValue
		fmt.Println("Plater: ", currentPlayer.name, "- goes -", currentPlayer.position)
		if Snake[currentPlayer.position] > 0 {
			currentPlayer.position = Snake[currentPlayer.position]
			fmt.Println("Player: ", currentPlayer.name, "- descends to -", currentPlayer.position)
		} else if Ladder[currentPlayer.position] > 0 {
			currentPlayer.position = Ladder[currentPlayer.position]
			fmt.Println("Player: ", currentPlayer.name, "- ascends to -", currentPlayer.position)
		}
		if currentPlayer.position > 85 {
			fmt.Println("Player: ~", currentPlayer.name, "~ Won")
		} else {
			Enqueue(currentPlayer)
		}
		
	}
}