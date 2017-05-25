package mcts

import (
	//"fmt"
	"github.com/pedroallenrevez/goTheory/gameTheory"
)

// State - provides the structure for a game state. which only needs the game
// being played
type State struct {
	Game gameTheory.Game
}

// StInterface - provides the interface for the state
type StInterface interface {
	Init()
	IsTerminal()
	Reward()
}

// Init - initializes a state
func (s *State) Init() {

}

// IsTerminal - checks if the state node is terminal
func (s *State) IsTerminal() bool {
	return false
}

// Reward - provides reward for a state
func (s *State) Reward() float64 {
	//check action
	return 0.0
}

//_____________________________________________________________________________
