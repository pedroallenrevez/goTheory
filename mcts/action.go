package mcts

//__________________________________ACTION______________________________________

// aID unique identifier for building actions
var aID = -1

// Action specification for an action
type Action struct {
	Name string
	ID   int
	Xcor int
	Ycor int
}

// AInterface you can only create an action
type AInterface interface {
	//CreateAction(string) *Action
}

// CreateAction creates an action with an unique identifier
// this action lacks context
func CreateAction(name string, x, y int) *Action {
	aID++
	new := Action{
		Name: name,
		ID:   aID,
		Xcor: x,
		Ycor: y,
	}
	return &new
}

//ApplyEffects - applies action to worldmodel
func (a Action) ApplyEffects(state *WorldModel) {
	//TODO moveplayer according to xcor and ycor
	MovePlayer(state.Map, state.SID, a.Xcor, a.Ycor)

}

//______________________________________________________________________________

/*
// GetAngles function responsible for creating angle actions
func (a Action) GetAngles(k int) []*Action {
	newActions := make([]*Action)
	turn0 := func(gman *GameManager) {
		gman.MovePlayer()
	}
	zero := ActInt.CreateAction("0", turnzero)

	turn45 := func(gman *GameManager) {
		gman.MovePlayer()
	}
	return newActions
}
*/
