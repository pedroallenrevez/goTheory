package mcts

// Cell struct responsible for the map
type Cell struct {
	X, Y   int
	Skater *Skater
}

// WorldMap - abstraction for the map with skaters
type WorldMap struct {
	//Do I want pointers in this?
	Map     []Cell
	Width   int
	Height  int
	Skaters []Skaters
}

// GetCell get cell for position
func (w *WorldMap) GetCell(x, y int) *Cell {
	//check torus behavior
	//TODO make it general
	//getcell being the main method to interact with the cells, we always
	//adjust the correct cell according to the torus behavior
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
	if index < 0 || index > 99 {
		fmt.Print("Bad result with ")
		fmt.Print(xcor, ycor)
		fmt.Print(" was: ")
		fmt.Print(x, y)
	}
	return &w.Map[index]
}

// GetSkater - get skater by id
func (w WorldMap) GetSkater(id int) *Skater {
	for _, skater := range w.Skaters {
		if skater.SID == id {
			return &skater
		}
	}
	return nil
}

// MovePlayer moveplayer from position 1 to pos2. torus functionality resides
// here
func (w *WorldMap) MovePlayer(sid int, action *Action) error {
	//TODO return error, and check if it was executed that way
	//TODO make general
	//if not  collision
	skater := GetSkater(sid)
	var errormsg error
	if !(w.checkCollision(skater, action)) {
		targetX := skater.PosX + action.Xcor
		targetY := skater.PosY + action.Ycor
		//set new skater position
		skater.PosX = targetX
		skater.PosY = targetY
		//assign skater to new possition
		targetCell := w.GetCell(targetX, targetY).Skater
		targetCell.Skater = skater
		//set old cell skater nil
		w.GetCell(skater.PosX, skater.PosY).Skater = nil

		errormsg = nil
	} else {
		//error = 1
		errormsg = errors.New("collision")

	}

	return errormsg
}

func (w WorldModel) checkCollision(skater *Skater, action *Action) bool {
	cell := w.GetCell(skater.PosX+action.Xcor, skater.PosY+action.Ycor)
	if cell.Skater != nil {
		return true
	}
	return false
}

// IsExecutable checks if action is executable (does not collide)
func (w WorldModel) IsExecutable(skater *Skater, action *Action) bool {
	return checkCollision(skater, action)
}
