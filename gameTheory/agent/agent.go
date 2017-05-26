package agent

import "github.com/pedroallenrevez/goTheory/gameTheory/action"

//__________________________________AGENT_______________________________________

// pID unique identifier for building agents
var pID = -1

var mixed = func() *action.Action {
	return nil
}
var alld = func() *action.Action {
	return nil
}
var allc = func() *action.Action {
	return nil
}
var tft = func() *action.Action {

	return nil
}

// Agent Specification for an agent
type Agent struct {
	Strat      Strategy
	Actions    []*action.Action
	ID         int
	TotalScore float64
}

//AgInterface interface responsible for creating an agent
type AgInterface interface {
	CreateAgent([]*action.Action, Strategy) *Agent
	CreateAgent2([]*action.Action) *Agent
	Init([]*action.Action, Strategy)
	Play() *action.Action
}

// Strategy A function to pick the action of all the available ones
// can be a mixed or pure strategy, even Tit for Tat, which adds a little more
// functionality
type Strategy func() *action.Action

// CreateAgent factory method for dumping agents with unique identifier
func (a Agent) CreateAgent(actions []*action.Action, strat Strategy) *Agent {
	pID++
	new := Agent{
		Strat:   strat,
		Actions: actions,
		/*tParam:     tVal,
		rParam:     rVal,
		pParam:     pVal,
		sParam:     sVal,*/
		TotalScore: 0,
		ID:         pID,
	}
	return &new

}

// Init initializes agent
func (a Agent) Init(actions []*action.Action, strat Strategy) {
	pID++
	a.ID = pID
	a.TotalScore = 0
	a.Actions = actions
	a.Strat = strat
}

// Play just returns the assigned strategy function
func (a Agent) Play() *action.Action {
	return a.Strat()
}
