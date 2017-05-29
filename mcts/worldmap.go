package mcts

import (
	"fmt"
	"strconv"
)

// Cell struct responsible for the map
type Cell struct {
	X, Y   int
	Skater *Skater
}

// WorldMap - abstraction for the map with skaters
type WorldMap struct {
	//Do I want pointers in this?
	Map     []Cell
	Skaters map[int]*Skater
	Width   int
	Height  int
}

// UpdateMap - sets the map for the
func (w *WorldMap) UpdateMap(mapeamento []Cell) {
	w.Map = mapeamento
}

/*
// SetSkaters - sets the map for the
func (w *WorldMap) SetSkaters(skaters []Skater) {
	w.Skaters = skaters
}
*/

// GetCell get cell for position
func (w *WorldMap) GetCell(x, y int) *Cell {
	//check torus behavior
	var xcor = x
	var ycor = y
	if x < 0 {
		xcor = w.Width - 1
	}
	if x >= w.Width {
		xcor = 0
	}
	if y < 0 {
		ycor = w.Height - 1
	}
	if y >= w.Height {
		ycor = 0
	}
	index := xcor*w.Width + ycor
	if index < 0 || index > w.Width*w.Height-1 {
		panic("bad cell")
	}
	return &w.Map[index]
}

// GetSkater - get skater by id
func (w WorldMap) GetSkater(id int) *Skater {
	return w.Skaters[id]
}

// AddSkater - add skater to map
func (w *WorldMap) AddSkater(skater *Skater) {
	cell := w.GetCell(skater.PosX, skater.PosY)
	if cell.Skater != nil {
		panic("there is askater there!")
	} else {
		//add to dictionary
		w.Skaters[skater.SID] = skater
		cell.Skater = skater
	}

}

// MovePlayer moveplayer from position 1 to pos2. torus functionality resides
// here
func (w *WorldMap) MovePlayer(sid int, action *Action) {
	skater := w.GetSkater(sid)
	if !(w.checkCollision(skater, action)) {
		targetX := skater.PosX + action.Xcor
		targetY := skater.PosY + action.Ycor
		//assign skater to new possition
		targetCell := w.GetCell(targetX, targetY)
		targetCell.Skater = skater
		targetCell.Skater.SetPos(targetCell.X, targetCell.Y)
		//set old cell skater nil
		w.GetCell(skater.PosX, skater.PosY).Skater = nil

	}
}

func (w *WorldMap) checkCollision(skater *Skater, action *Action) bool {
	cell := w.GetCell(skater.PosX+action.Xcor, skater.PosY+action.Ycor)
	if cell.Skater != nil {
		return true
	}
	return false
}

// IsExecutable checks if action is executable (does not collide)
func (w *WorldMap) IsExecutable(sid int, action *Action) bool {
	return w.checkCollision(w.GetSkater(sid), action)
}

//CopyMap - creates a copy of the current world map for the worldModel
func (w WorldMap) CopyMap() WorldMap {
	newMap := WorldMap{
		Map:     w.Map,
		Width:   w.Width,
		Height:  w.Height,
		Skaters: w.Skaters,
	}
	return newMap
}

// PrintGame - prints current state of the board
func (w WorldMap) PrintGame() {
	//map
	for i := 0; i < w.Width; i++ {
		fmt.Print("|")
		for j := 0; j < w.Height; j++ {
			cell := w.GetCell(i, j)
			if cell.Skater != nil {
				fmt.Print(" " + strconv.Itoa(cell.Skater.SID) + " " + "|")
				continue
			} else {
				fmt.Print("   " + "|")
			}

		}
		fmt.Println()
	}
	fmt.Println()

}

// PrintSkaters - prints all skater positions
func (w WorldMap) PrintSkaters() {
	fmt.Println()
	fmt.Println("Skaters")
	for _, skater := range w.Skaters {
		w.PrintSkater(skater.SID)
	}
	fmt.Println()

}

// PrintSkater - prints a single skater
func (w WorldMap) PrintSkater(sid int) {
	skater := w.GetSkater(sid)

	fmt.Print(" x: ")
	fmt.Print(skater.PosX)
	fmt.Print(" y: ")
	fmt.Print(skater.PosY)
	fmt.Print(" direction: ")
	fmt.Print(skater.Direction.Name)
	skater.Mcts.PrintMcts()
	fmt.Println()
}

// PrintMap - prints thw whole thing
func (w WorldMap) PrintMap() {
	fmt.Println(w.Width, w.Height)
	w.PrintGame()
	w.PrintSkaters()
}
