package main

import "fmt"

func main() {
	players := GetPlayers()
	QueuePlayers(players)
	result := StartGame()
	for position, player := range result {
		if player != 0 {
			fmt.Printf("Player %d finished at position %d\n", player, position)
		}
	}
}