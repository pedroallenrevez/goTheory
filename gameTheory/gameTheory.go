package gameTheory

import (
	"fmt"
)

// pID unique identifier for building agents
var pID = -1

// aID unique identifier for building actions
var aID = -1

//__________________________________ACTION______________________________________

// Action specification for an action
type Action struct {
	name string
	id   int
}

// CreateAction creates an action with an unique identifier
// this action lacks context
func (a Action) CreateAction(name string) *Action {
	id = aID + 1
	new := Action{
		name: name,
		id:   id,
	}
	return &new
}

//__________________________________AGENT_______________________________________

// Agent Specification for an agent
type Agent struct {
	Strat      Strategy
	Actions    []Action
	id         int
	totalScore float64
	//PD values
	tParam, rParam, sParam, pParam float64
}

// Strategy A function to pick the action of all the available ones
// can be a mixed or pure strategy, even Tit for Tat, which adds a little more
// functionality
type Strategy func() Action

// CreateAgent factory method for dumping agents with unique identifier
func (a Agent) CreateAgent(actions []Action, strat Strategy, tVal, rVal, pVal, sVal float64) *Agent {
	id := pID + 1
	new := Agent{
		Strat:      strat,
		Actions:    actions,
		tParam:     tVal,
		rParam:     rVal,
		pParam:     pVal,
		sParam:     sVal,
		totalScore: 0,
		id:         id,
	}
	return &new

}

// PickAction ????
func (a *Agent) PickAction() Action {

}

//__________________________________CORE________________________________________

// Game structure for a normal n x n games
type Game struct {
	Actions []Action
	Players []Agent
	//only has a dictionary of dictionary of states

	//TODO implement nmap with n players
	//2 player implementation
}

type state struct {
	//payoffs are ordered according to player ID
	payoffs []float64
}

// CreateGame creates a game with n actions, m players
func (g *Game) CreateGame(actions []Action, players []Agent) Game {
	// maps of maps dictionaries
	// 2x2 prisoners dilemma
	// do nxn games and 2x2 games. 1on1 or N on N

}

// CalculatePayoff returns payoff for given game players and action
// it is not needed to compute all the values at the beggining, because the
// choice of action is deterministic. We just calculate it when needed.
func (g *Game) CalculatePayoff(g Game, p Agent, act Action) {
	//run q function
}

// Solve returns payoff for each agent in game
func (g Game) Solve(g Game) float64 {
	//add score to each participating agent
	//for all agents - calculate payoff
}

// LoadPrisonerDilemma receive a game and assign the PD values to them
func (g *Game) LoadPrisonerDilemma(g Game) Game {

}

//______________________________________________________________________________

// FindNash finds all NE present in the game
// ret value?
func (g *Game) FindNash(g Game) {

}

// FindParetto finds all NE present in the game
func (g *Game) FindParetto(g Game) {

}
