package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func QueuePlayers(players int) {
	for i:=1; i <= players; i++ {
		Enqueue(Player{
			name: i,
			position: 0,
		})
	}
}

func GetPlayers() int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Enter number of players (1-4): ")
		Input, _ := reader.ReadString('\n')
		playerCount, _ := strconv.ParseFloat(strings.TrimSpace(Input), 64)
		fmt.Println(playerCount)
		if playerCount > 0 && playerCount < 5 {
			return int(playerCount)
		}
	}
}