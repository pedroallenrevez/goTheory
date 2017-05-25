package sk8rBoy

import (
	"github.com/pedroallenrevez/goTheory/gameTheory/action"
	"github.com/pedroallenrevez/goTheory/gameTheory/agent"
)

// Cell struct responsible for the map
type Cell struct {
	X, Y int
	sk8r Skater
}

// Skater struct that embeds the agent behaviour
type Skater struct {
	agent.Agent //includes ID
	Direction   *action.Action
}

// Init - initialize all the actions
func (s *Skater) CreateSkater(actions []action.Action) {
	//create actions here and assign to the agent
	//call superclass method for agent and set direction

	newAgent := CreateAgent(actions, strat)

}

// GameManager instance responsible for the running the game
type GameManager struct {
	Map             [][]Cell
	AngleResolution int //ex.: 60 (0, 60, 120, 180, 240, 300
	Speed           int //displacement - 1sq or 2sq
	CollisionRadius int //1sq or 2sq
	RewardExec      float64
	RewardNotExec   float64
}

// Init create a sk8rboy game with the specified parameters
func (gman GameManager) Init(w, h int, k int, speed int, radius int, r1, r2 float64) {
	//add sk8rs to array
	//init map with size w,h
	newMap := make
	//create actions for k angles
	new := GameManager{
		Map:             newMap,
		AngleResolution: k,
		Speed:           speed,
		CollisionRadius: radius,
		RewardExec:      r1,
		RewardNotExec:   r2,
	}
}

// MovePlayers moveplayer from position 1 to pos2. torus functionality resides
// here
func (gman GameManager) MovePlayer(x1, y1, x2, y2 int) {
	//if out of bound go to other side
	//do remainder calculus for going off board
}

func (gman GameManager) Update() {
	//update all sk8rs one by one
}
