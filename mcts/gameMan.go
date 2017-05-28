package mcts

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var global *GameManager
var sID = -1

// Skater struct that embeds the agent behaviour
type Skater struct {
	Direction *Action
	PosX      int
	PosY      int
	SID       int
}

// GameManager instance responsible for the running the game
type GameManager struct {
	Map             WorldMap
	Speed           int //displacement - 1sq or 2sq
	CollisionRadius int //1sq or 2sq not needed
	RewardExec      float64
	RewardNotExec   float64
	Actions         []*Action
}

// GManInterface provides the accessible methods for the game manager
type GManInterface interface {
	Init(int, int, int, int, int, float64, float64)
	MovePlayer(int, int, int, int)
	PrintGame()
	PrintSkaters()
	Update()
}

// Init create a sk8rboy game with the specified parameters
func (gman *GameManager) Init(w, h int, n int, speed int, radius int, r1, r2 float64) {
	global = gman
	gman.Speed = speed
	gman.RewardExec = r1
	gman.RewardNotExec = r2
	gman.CollisionRadius = radius
	gman.Width = w
	gman.Height = h
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
	gman.Map = newMap
	/*
		for i := 0; i < gman.Width; i++ {
			for j := 0; j < gman.Height; j++ {
				fmt.Println(gman.GetCell(i, j).X, gman.GetCell(i, j).Y)
			}
		}
	*/
	gman.CreateActions()
	gman.InitAgents(n, gman.Actions)

}

// InitAgents - initializes n agents
func (gman *GameManager) InitAgents(n int, actions []*Action) {
	//add sk8rs to array
	skaters := make([]*Skater, n)
	for i := 0; i < n; i++ {
		newAgent := CreateSkater()
		skaters[i] = newAgent
	}
	gman.Skaters = skaters
}

// CreateSkater - initialize all the actions
func CreateSkater() *Skater {
	sID++
	newSk8r := &Skater{
		Direction: global.Actions[rand.Intn(len(global.Actions))],
		PosX:      rand.Intn(global.Width),
		PosY:      rand.Intn(global.Height),
		SID:       sID,
	}

	//add them to the gameManagerMap
	global.GetCell(newSk8r.PosX, newSk8r.PosY).Skater = newSk8r
	return newSk8r

}

// CreateActions - create a number of actions according to angle resolution
func (gman *GameManager) CreateActions() {
	//do 8 angles
	actions := make([]*Action, 8)
	turn0 := CreateAction("0", 1, 0)
	actions[0] = turn0
	turn45 := CreateAction("45", 1, -1)
	actions[1] = turn45
	turn90 := CreateAction("90", 0, -1)
	actions[2] = turn90
	turn135 := CreateAction("135", -1, -1)
	actions[3] = turn135
	turn180 := CreateAction("180", -1, 0)
	actions[4] = turn180
	turn225 := CreateAction("225", -1, 1)
	actions[5] = turn225
	turn270 := CreateAction("270", 0, 1)
	actions[6] = turn270
	turn315 := CreateAction("315", 1, 1)
	actions[7] = turn315

	gman.Actions = actions

}

// Update - updates the game, all the skaters, and plays all games
func (gman *GameManager) Update() {
	//update all sk8rs one by one
	for _, sk8r := range gman.Skaters {
		//run mcts
		nextAction := gman.Actions[rand.Intn(len(gman.Actions))]
		sk8r.Direction = nextAction

		gman.Map.MovePlayer(sk8r.SID, nextAction)
	}
}

//______________________________________________________________________________

//CopyMap - creates a copy of the current world map for the worldModel
func (gman *GameManager) CopyMap() WorldMap {
	return gman.Map
}

// PrintGame - prints current state of the board
func (gman *GameManager) PrintGame() {
	//map
	for i := 0; i < gman.Width; i++ {
		for j := 0; j < gman.Height; j++ {
			cell := gman.GetCell(i, j)
			if cell.Skater != nil {
				fmt.Print(" " + strconv.Itoa(cell.Skater.ID) + " ")
				continue
			} else {
				fmt.Print(" _ ")
			}

		}
		fmt.Println()
	}
}

// PrintSkaters - prints all skater positions
func (gman *GameManager) PrintSkaters() {
	for _, skater := range gman.Skaters {
		fmt.Print(" x: ")
		fmt.Print(skater.PosX)
		fmt.Print(" y: ")
		fmt.Print(skater.PosY)
		fmt.Print(" direction: ")
		fmt.Print(skater.Direction.Name)
		fmt.Println()
	}

}
