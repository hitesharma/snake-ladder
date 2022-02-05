package main

import (
	"errors"
)

// init players queue
var queue []Player

// append player to queue end
func Enqueue(element Player) {
	queue = append(queue, element)
}

// slice off player from queue start
func Dequeue() (Player, error) {
	if len(queue) == 0 {
		return Player{}, errors.New("queue empty")
	}
	element:=queue[0]
	queue = queue[1:]
	return element, nil
}
