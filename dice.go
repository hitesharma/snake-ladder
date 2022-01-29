package main

import (
	"math/rand"
	"time"
)

// generate random value b/w 1 & 6
func RollDice() int {
	rand.Seed(time.Now().UnixNano())
	a:=rand.Intn(6-1)+1
	return a
}