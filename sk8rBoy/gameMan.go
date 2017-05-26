package sk8rBoy

import (
	"github.com/pedroallenrevez/goTheory/gameTheory/action"
	"github.com/pedroallenrevez/goTheory/gameTheory/agent"
	"github.com/pedroallenrevez/goTheory/mcts"
	"math/rand"
	"time"
)

// Cell struct responsible for the map
type Cell struct {
	X, Y int
	Sk8r *Skater
}

// Skater struct that embeds the agent behaviour
type Skater struct {
	Direction *action.Action
	PosX      int
	PosY      int
}

// CreateSkater - initialize all the actions
func (s *Skater) CreateSkater() {
	//create actions here and assign to the agent
	//call superclass method for agent and set direction

	//set random direction and pos
	//add them to the gameManagerMap

}

// GameManager instance responsible for the running the game
type GameManager struct {
	Map             []*Cell
	Width           int
	Height          int
	Speed           int //displacement - 1sq or 2sq
	CollisionRadius int //1sq or 2sq not needed
	RewardExec      float64
	RewardNotExec   float64
	Skaters         []*Skater
	Actions         []*action.Action
}

// GManInterface provides the accessible methods for the game manager
type GManInterface interface {
	Init(int, int, int, int, int, float64, float64)
	MovePlayer(int, int, int, int)
	Update()
}

// Init create a sk8rboy game with the specified parameters
func (gman GameManager) Init(w, h int, n int, speed int, radius int, r1, r2 float64) {
	rand.Seed(time.Now().Unix())
	//init map with size w,h
	newMap := make([]*Cell, w*h) //access with [i*m + j]
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			newCell := &Cell{
				X: i,
				Y: j,
			}
			newMap[i*w+j] = newCell
		}
	}
	gman.InitAgents(n)
	gman.CreateActions()

	gman.Speed = speed
	gman.RewardExec = r1
	gman.RewardNotExec = r2
	gman.CollisionRadius = radius
}

// InitAgents - initializes n agents
func (gman *GameManager) InitAgents(n int, actions []*action.Action) {
	//add sk8rs to array
	skaters := make([]*Skater)
	var ag = new(agent.Agent)
	var AgInt agent.AgInterface = ag
	for i := range n {
		newAgent := ag.CreateSkater(actions)
		skaters = append(skaters, newAgent)
	}
	gman.Skaters = skaters
}

// CreateActions - create a number of actions according to angle resolution
func (gman *GameManager) CreateActions() {
	//do 8 angles
	actions := make([]*action.Action)
	var act = new(action.Action)
	var ActInt action.AInterface = act
	turn0 := act.CreateAction("0")
	turn45 := act.CreateAction("45")
	turn90 := act.CreateAction("90")
	turn135 := act.CreateAction("135")
	turn180 := act.CreateAction("180")
	turn225 := act.CreateAction("225")
	turn270 := act.CreateAction("270")
	turn315 := act.CreateAction("315")
	actions := append(actions, turn0)
	actions := append(actions, turn45)
	actions := append(actions, turn90)
	actions := append(actions, turn135)
	actions := append(actions, turn180)
	actions := append(actions, turn225)
	actions := append(actions, turn270)
	actions := append(actions, turn315)
	gman.Actions = actions

}

// MovePlayer moveplayer from position 1 to pos2. torus functionality resides
// here
func (gman *GameManager) MovePlayer(x1, y1, x2, y2 int) {
	//check collision for safety?
	gman.GetCell(x2, y2).Sk8r = gman.GetCell(x1, y1).Sk8r
	gman.GetCell(x1, y1).Sk8r = nil
	gman.GetCell(x2, y2).Sk8r.PosX = x2
	gman.GetCell(x2, y2).Sk8r.PosY = y2
}

// GetCell get cell for position
func (gman *GameManager) GetCell(x, y int) *Cell {
	//check torus behavior
	//getcell being the main method to interact with the cells, we always
	//adjust the correct cell according to the torus behavior
	var xcor = x
	var ycor = y
	if x < 0 {
		xcor = gman.Width - 1 - gman.Speed
	}
	if x > gman.Width {
		xcor = gman.Speed - 1
	}
	if y < 0 {
		ycor = gman.Heigth - 1 - gman.Speed
	}
	if y > gman.Height {
		ycor = gman.Speed - 1
	}
	return gman.Map[xcor*gman.Width+ycor]
}

// Update - updates the game, all the skaters, and plays all games
func (gman *GameManager) Update() {
	//update all sk8rs one by one
	for sk8r := range gman.Skaters {
		//run mcts
		//Horrible hack pedro
		//TODO fix actions being hardcoded. actions should apply effects
		//to state
		nextAction := gman.Actions[rand.Intn(len(gman.Actions))]

		//check for collision should be made on state?
		if nextAction.name == "0" {
			gman.MovePlayer(sk8r.PosX, sk8r.PosY, sk8r.PosX+1, sk8r.PosY)
		}
		if nextAction.name == "45" {
			gman.MovePlayer(sk8r.PosX, sk8r.PosY, sk8r.PosX+1, sk8r.PosY+1)
		}
		if nextAction.name == "90" {
			gman.MovePlayer(sk8r.PosX, sk8r.PosY, sk8r.PosX, sk8r.PosY+1)
		}
		if nextAction.name == "135" {
			gman.MovePlayer(sk8r.PosX, sk8r.PosY, sk8r.PosX-1, sk8r.PosY+1)
		}
		if nextAction.name == "180" {
			gman.MovePlayer(sk8r.PosX, sk8r.PosY, sk8r.PosX-1, sk8r.PosY)
		}
		if nextAction.name == "225" {
			gman.MovePlayer(sk8r.PosX, sk8r.PosY, sk8r.PosX-1, sk8r.PosY-1)
		}
		if nextAction.name == "270" {
			gman.MovePlayer(sk8r.PosX, sk8r.PosY, sk8r.PosX, sk8r.PosY-1)
		}
		if nextAction.name == "315" {
			gman.MovePlayer(sk8r.PosX, sk8r.PosY, sk8r.PosX+1, sk8r.PosY-1)
		}
		//check collision - this is in state
		//do action - move player
	}
}
