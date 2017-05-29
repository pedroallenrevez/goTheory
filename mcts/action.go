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

//ApplyEffects - applies action to worldmodel, generates a child with reward
func (a Action) ApplyEffects(state WorldModel) WorldModel {
	var reward float64
	if state.Map.IsExecutable(state.SID, &a) {
		reward = globalReward1
	} else {
		reward = globalReward2
	}
	newState := state.GenerateChild(reward, &a)
	newState.Map.MovePlayer(state.SID, &a)
	return newState
}
