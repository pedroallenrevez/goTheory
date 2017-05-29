package mcts

import (
	"fmt"
	"math/rand"
	"time"
)

var globalReward1 float64
var globalReward2 float64

var globalWorldMap *WorldMap
var globalActions []*Action
var globalSkaters []*Skater

var sID = -1

// Skater struct that embeds the agent behaviour
type Skater struct {
	Direction *Action
	PosX      int
	PosY      int
	SID       int
	Mcts      *MCTS
}

//SetMcts - set the search algorithm for the agent
func (s *Skater) SetMcts(mcts *MCTS) {
	s.Mcts = mcts
}

//SetPos - set the search algorithm for the agent
func (s *Skater) SetPos(x, y int) {
	s.PosX = x
	s.PosY = y
}

// CreateSkater - initialize all the actions
func (gman *GameManager) CreateSkater() {
	/*
		sID++
		newSk8r := Skater{
			Direction: gman.RandomAction(),
			PosX:      rand.Intn(gman.Width - 1),
			PosY:      rand.Intn(gman.Height - 1),
			SID:       sID,
		}

		//add them to the gameManagerMap
		return newSk8r
	*/

}

// GameManager instance responsible for the running the game
type GameManager struct {
	Map             *WorldMap
	Speed           int //displacement - 1sq or 2sq
	CollisionRadius int //1sq or 2sq not needed
}

// GManInterface provides the accessible methods for the game manager
type GManInterface interface {
	Init(int, int, int, int, int, float64, float64)
	PrintSkaters()
	PrintGame()
	Update()
}

// Init create a sk8rboy game with the specified parameters
func (gman *GameManager) Init(w, h int, n int, speed int, radius int, r1, r2 float64) {
	globalReward1 = r1
	globalReward2 = r2
	gman.Speed = speed
	gman.CollisionRadius = radius
	rand.Seed(time.Now().Unix())

	//init actions
	globalActions = gman.InitActions()
	//Init agents
	globalSkaters = gman.InitAgents(n, w, h)
	//Init WorldMap
	globalWorldMap = gman.InitWorldMap(w, h)
	gman.Map = globalWorldMap
	//init mcts
	gman.InitMcts()

}

// InitActions - create a number of actions according to angle resolution
func (gman *GameManager) InitActions() []*Action {
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
	return actions
}

// InitAgents - initializes n agents
func (gman *GameManager) InitAgents(n, w, h int) []*Skater {
	//add sk8rs to array
	skaters := make([]*Skater, 0)
	for i := 0; i < n; i++ {
		//build state
		//build node
		/*
		 */
		//create agent with mcts
		sID++
		posX := rand.Intn(w - 1)
		posY := rand.Intn(h - 1)
		for _, skater := range skaters {
			if skater.PosX == posX {
				posX = rand.Intn(w - 1)
			}
			if skater.PosY == posY {
				posY = rand.Intn(h - 1)
			}
		}
		newAgent := &Skater{
			Direction: RandomAction(),
			PosX:      posX,
			PosY:      posY,
			SID:       sID,
		}
		skaters = append(skaters, newAgent)
	}

	return skaters
}

// InitWorldMap - initializes n agents
func (gman *GameManager) InitWorldMap(w, h int) *WorldMap {
	newMap := make([]Cell, w*h) //access with [i*m + j]
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			var newCell Cell
			newCell = Cell{
				X: i,
				Y: j,
			}
			newMap[i*w+j] = newCell
		}
	}
	//worldmodel
	newWorldMap := &WorldMap{
		Map:     newMap,
		Skaters: make(map[int]*Skater),
		Width:   w,
		Height:  h,
	}
	for _, skater := range globalSkaters {
		//on add skater fill the map with sid
		newWorldMap.AddSkater(skater)
	}
	return newWorldMap

}

// InitMcts - initializes n agents
func (gman *GameManager) InitMcts() {
	for _, skater := range globalWorldMap.Skaters {
		newState := WorldModel{
			SID:       skater.SID,
			Fitness:   0,
			TurnsLeft: 100,
			Actions:   make([]*Action, 0),
			Reward:    0,
			Map:       *gman.Map,
		}
		root := &Node{
			CurrentState:   newState,
			Parent:         nil,
			ChildNodes:     make([]*Node, 0),
			Reward:         0,
			UntriedActions: globalActions,
		}
		//Create Mcts
		newMCTS := &MCTS{
			C:             1.4,
			CurrentState:  newState,
			Root:          root,
			ExpandedNodes: make([]*Node, 0),
		}
		skater.SetMcts(newMCTS)
	}
}

// Update - updates the game, all the skaters
func (gman *GameManager) Update() {
	//update all sk8rs one by one
	for _, sk8r := range globalWorldMap.Skaters {
		//run mcts
		//nextAction := gman.RandomAction()
		//go routine and wait for model on channel at end of update?
		nextAction := sk8r.Mcts.Run()
		sk8r.Direction = nextAction
		//failsafe in case there is not action?
		globalWorldMap.MovePlayer(sk8r.SID, nextAction)
	}

	for _, sk8r := range globalWorldMap.Skaters {
		fmt.Println("updated maps")
		sk8r.Mcts.UpdateExpandedMaps(*globalWorldMap)
		sk8r.Mcts.Reset()
	}
}

//______________________________________________________________________________

// RandomAction - returns a random action
func RandomAction() *Action {
	return globalActions[rand.Intn(len(globalActions))]
}

// PrintGameMan - calls game manager map method for printing
func PrintGameMan(gman *GameManager) {
	gman.PrintGame()
	gman.PrintSkaters()
}

// PrintGame - calls game manager map method for printing
func (gman *GameManager) PrintGame() {
	gman.Map.PrintGame()
}

// PrintSkaters - prints all skaters
func (gman *GameManager) PrintSkaters() {
	fmt.Print()
	gman.Map.PrintSkaters()
}
