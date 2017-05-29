package mcts

import (
	"fmt"
)

// WorldModel - worldmodel abstraction for sk8rboy
type WorldModel struct {
	//score
	SID       int
	Map       WorldMap
	Fitness   float64
	Reward    float64
	TurnsLeft int
	Actions   []*Action
}

// IsTerminal - checks if node is terminal or not
func (w *WorldModel) IsTerminal() bool {
	if w.TurnsLeft == 0 {
		return true
	}
	return false

}

// GenerateChild - gens a child world model
func (w *WorldModel) GenerateChild(reward float64, action *Action) WorldModel {
	newState := WorldModel{
		SID:       w.SID,
		Fitness:   w.Fitness + reward, //fitness + action reward
		TurnsLeft: w.TurnsLeft - 1,
		Actions:   append(w.Actions, action),
		Reward:    reward,
		Map:       w.Map.CopyMap(),
	}
	return newState
}

// UpdateMap - prints the state in a fashionly order
func (w *WorldModel) UpdateMap(mapa WorldMap) {
	w.Map = mapa
}

// PrintState - prints the state in a fashionly order
func (w *WorldModel) PrintState() {
	fmt.Print("SkaterID: ")
	fmt.Print(w.SID)
	fmt.Println("Map:")
	w.Map.PrintGame()
	fmt.Print("Fitness: ")
	fmt.Print(w.Fitness)
	fmt.Print("Reward: ")
	fmt.Print(w.Reward)
	fmt.Print("TurnsLeft: ")
	fmt.Print(w.TurnsLeft)
	w.Map.PrintSkaters()

}
