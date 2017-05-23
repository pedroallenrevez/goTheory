package agent

import "github.com/pedroallenrevez/goTheory/gameTheory/action"

//__________________________________AGENT_______________________________________

// pID unique identifier for building agents
var pID = -1

// Agent Specification for an agent
type Agent struct {
	Strat      Strategy
	Actions    []*action.Action
	id         int
	totalScore float64
	//PD values
}

//AgInterface interface responsible for creating an agent
type AgInterface interface {
	CreateAgent([]*action.Action, Strategy) *Agent
}

// Strategy A function to pick the action of all the available ones
// can be a mixed or pure strategy, even Tit for Tat, which adds a little more
// functionality
type Strategy func() *action.Action

// CreateAgent factory method for dumping agents with unique identifier
func (a Agent) CreateAgent(actions []*action.Action, strat Strategy /*, tVal, rVal, pVal, sVal float64*/) *Agent {
	pID++
	new := Agent{
		Strat:   strat,
		Actions: actions,
		/*tParam:     tVal,
		rParam:     rVal,
		pParam:     pVal,
		sParam:     sVal,*/
		totalScore: 0,
		id:         pID,
	}
	return &new

}
