package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

/*
1.
 The game will ask the player how many tries he will have before losing the game
2.
 The number to guess must be between 0 and 10.
3.
 Every time a player enters a number to guess, the number of retries drops by one.
4.
 If the number of retries reaches zero and the number is still incorrect the game
finishes and the player has lost.
5. If the player guess the number, he wins.
*/

// GameContext and a game context to store
// the information between states.
// the context needs to store the number of retries,
// if the use has won or not,
// the secret number to guess
// and the current state
type GameContext struct {
	SecretNum int
	Retries   int
	Won       bool
	Next      GameState
}

// GameState we need the interface to represent the different states
type GameState interface {
	executeState(*GameContext) bool
}

type StartState struct {
}

type AskState struct {
}

type WinState struct{}

func (w *WinState) executeState(c *GameContext) bool {
	println("Congrats, you won")
	return false
}

type LoseState struct{}

func (l *LoseState) executeState(c *GameContext) bool {
	fmt.Printf("You lose. The correct number was: %d\n", c.SecretNum)
	return false
}

func (a *AskState) executeState(c *GameContext) bool {
	fmt.Printf("Introduce a number between 0 and 10, you have %d tries left\n", c.Retries)

	var n int
	fmt.Fscan(os.Stdin, &n)
	c.Retries--
	if n == c.SecretNum {
		c.Won = true
		c.Next = &WinState{}
	}
	if c.Retries == 0 {
		c.Next = &LoseState{}
	}
	return true
}

func (s *StartState) executeState(c *GameContext) bool {
	c.Next = &AskState{}
	rand.Seed(time.Now().UnixNano())
	c.SecretNum = rand.Intn(10)
	fmt.Println("Introduce a number a number of retries to set the difficulty:")
	fmt.Fscan(os.Stdin, &c.Retries)

	//  we return true to tell the “engine” that the game must continue.
	return true
}

func main() {
	start := StartState{}
	game := GameContext{Next: &start}
	for game.Next.executeState(&game) {
	}
}
