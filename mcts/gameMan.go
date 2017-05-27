package mcts

import (
	"math/rand"
	"time"
)

var global *GameManager

// Cell struct responsible for the map
type Cell struct {
	X, Y   int
	Skater *Skater
}

// Skater struct that embeds the agent behaviour
type Skater struct {
	Direction *Action
	PosX      int
	PosY      int
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
	Actions         []*Action
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
	gman.CreateActions()
	gman.InitAgents(n, gman.Actions)

	gman.Speed = speed
	gman.RewardExec = r1
	gman.RewardNotExec = r2
	gman.CollisionRadius = radius
	global = &gman
}

// InitAgents - initializes n agents
func (gman *GameManager) InitAgents(n int, actions []*Action) {
	//add sk8rs to array
	skaters := make([]*Skater, n)
	for i := 0; i < n; i++ {
		newAgent := CreateSkater()
		skaters = append(skaters, newAgent)
	}
	gman.Skaters = skaters
}

// CreateSkater - initialize all the actions
func CreateSkater() *Skater {
	newSk8r := &Skater{
		Direction: global.Actions[rand.Intn(len(global.Actions))],
		PosX:      rand.Intn(global.Width),
		PosY:      rand.Intn(global.Height),
	}

	//add them to the gameManagerMap
	global.GetCell(newSk8r.PosX, newSk8r.PosY).Skater = newSk8r
	return newSk8r

}

// CreateActions - create a number of actions according to angle resolution
func (gman *GameManager) CreateActions() {
	//do 8 angles
	actions := make([]*Action, 8)
	//this is stupid. CreateAction should be public
	turn0 := CreateAction("0")
	turn45 := CreateAction("45")
	turn90 := CreateAction("90")
	turn135 := CreateAction("135")
	turn180 := CreateAction("180")
	turn225 := CreateAction("225")
	turn270 := CreateAction("270")
	turn315 := CreateAction("315")
	actions = append(actions, turn0)
	actions = append(actions, turn45)
	actions = append(actions, turn90)
	actions = append(actions, turn135)
	actions = append(actions, turn180)
	actions = append(actions, turn225)
	actions = append(actions, turn270)
	actions = append(actions, turn315)
	gman.Actions = actions

}

// MovePlayer moveplayer from position 1 to pos2. torus functionality resides
// here
func (gman *GameManager) MovePlayer(x1, y1, x2, y2 int) {
	//check collision for safety?
	gman.GetCell(x2, y2).Skater = gman.GetCell(x1, y1).Skater
	gman.GetCell(x1, y1).Skater = nil
	gman.GetCell(x2, y2).Skater.PosX = x2
	gman.GetCell(x2, y2).Skater.PosY = y2
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
		ycor = gman.Height - 1 - gman.Speed
	}
	if y > gman.Height {
		ycor = gman.Speed - 1
	}
	return gman.Map[xcor*gman.Width+ycor]
}

// Update - updates the game, all the skaters, and plays all games
func (gman *GameManager) Update() {
	//update all sk8rs one by one
	for _, sk8r := range gman.Skaters {
		//run mcts
		//Horrible hack pedro
		//TODO fix actions being hardcoded. actions should apply effects
		//to state
		nextAction := gman.Actions[rand.Intn(len(gman.Actions))]

		//check for collision should be made on state?
		if nextAction.Name == "0" {
			gman.MovePlayer(sk8r.PosX, sk8r.PosY, sk8r.PosX+1, sk8r.PosY)
		}
		if nextAction.Name == "45" {
			gman.MovePlayer(sk8r.PosX, sk8r.PosY, sk8r.PosX+1, sk8r.PosY+1)
		}
		if nextAction.Name == "90" {
			gman.MovePlayer(sk8r.PosX, sk8r.PosY, sk8r.PosX, sk8r.PosY+1)
		}
		if nextAction.Name == "135" {
			gman.MovePlayer(sk8r.PosX, sk8r.PosY, sk8r.PosX-1, sk8r.PosY+1)
		}
		if nextAction.Name == "180" {
			gman.MovePlayer(sk8r.PosX, sk8r.PosY, sk8r.PosX-1, sk8r.PosY)
		}
		if nextAction.Name == "225" {
			gman.MovePlayer(sk8r.PosX, sk8r.PosY, sk8r.PosX-1, sk8r.PosY-1)
		}
		if nextAction.Name == "270" {
			gman.MovePlayer(sk8r.PosX, sk8r.PosY, sk8r.PosX, sk8r.PosY-1)
		}
		if nextAction.Name == "315" {
			gman.MovePlayer(sk8r.PosX, sk8r.PosY, sk8r.PosX+1, sk8r.PosY-1)
		}
		//check collision - this is in state
		//do action - move player
	}
}
