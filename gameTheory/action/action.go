package action

//__________________________________ACTION______________________________________

// aID unique identifier for building actions
var aID = -1

// Action specification for an action
type Action struct {
	name string
	id   int
}

// AInterface you can only create an action
type AInterface interface {
	CreateAction(string) *Action
}

// CreateAction creates an action with an unique identifier
// this action lacks context
func (a Action) CreateAction(name string) *Action {
	aID++
	new := Action{
		name: name,
		id:   aID,
	}
	return &new
}
