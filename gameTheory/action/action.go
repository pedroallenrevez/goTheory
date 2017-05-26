package action

//__________________________________ACTION______________________________________

// aID unique identifier for building actions
var aID = -1

// Action specification for an action
type Action struct {
	Name         string
	ID           int
	ApplyEffects func()
}

// AInterface you can only create an action
type AInterface interface {
	CreateAction(string) *Action
}

// CreateAction creates an action with an unique identifier
// this action lacks context
func (a Action) CreateAction(name string, effects func()) *Action {
	aID++
	new := Action{
		Name:         name,
		ID:           aID,
		ApplyEffects: effects,
	}
	return &new
}

//func (a Action) ApplyEffects(state WorldModel) {}

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
