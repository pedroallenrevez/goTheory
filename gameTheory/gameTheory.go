package gameTheory

import (
	"fmt"
	//	"strconv"
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

//__________________________________AGENT_______________________________________

// Agent Specification for an agent
type Agent struct {
	Strat      Strategy
	Actions    []*Action
	id         int
	totalScore float64
	//PD values
}

//AgInterface interface responsible for creating an agent
type AgInterface interface {
	CreateAgent([]*Action, Strategy) *Agent
}

// Strategy A function to pick the action of all the available ones
// can be a mixed or pure strategy, even Tit for Tat, which adds a little more
// functionality
type Strategy func() *Action

// CreateAgent factory method for dumping agents with unique identifier
func (a Agent) CreateAgent(actions []*Action, strat Strategy /*, tVal, rVal, pVal, sVal float64*/) *Agent {
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

// PickAction ????
func (a *Agent) PickAction() {
}

//__________________________________CORE________________________________________

// Game structure for a normal n x n games
type Game struct {
	Actions []*Action
	Players []*Agent
	//only has a dictionary of dictionary of states
	//TODO implement nmap with n players
	//2 player implementation
	States [][]state
}

// GInterface interfface responsible for the game method available
type GInterface interface {
	CreateGame([]Action, []Agent) *Game
	Solve(Game)
}

type state struct {
	//payoffs are ordered according to player ID
	payoffs []float64
}

// CreateGame creates a game with n actions, m players
func (g *Game) CreateGame(actions []*Action, players []*Agent) *Game {
	// maps of maps dictionaries
	// 2x2 prisoners dilemma
	// 1x1 or NxN

	//init states with size of actions we want a.len x a.len

	NIDstates := make([][]state, len(actions))

	for i := range NIDstates {
		NIDstates[i] = make([]state, len(actions))
	}

	payoffs := make([]float64, len(players), len(players))
	payoffs[0] = 0.0
	payoffs[1] = 0.0

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			NIDstates[i][j].payoffs = payoffs
		}
	}

	new := Game{
		Actions: actions,
		Players: players,
		States:  NIDstates,
	}
	return &new
}

// CalculatePayoff returns payoff for given game players and action
// it is not needed to compute all the values at the beggining, because the
// choice of action is deterministic. We just calculate it when needed.
func (g *Game) CalculatePayoff(game *Game, p *Agent, act *Action) {
	//run q function
}

// Solve returns payoff for each agent in game
func (g Game) Solve(game *Game) float64 {
	//add score to each participating agent
	//for all agents - calculate payoff
	return 0.0
}

/*
// LoadPrisonerDilemma receive a game and assign the PD values to them
// THIS ONLY WILL WORK FOR 2x2 GAMES
func (g *Game) LoadPrisonerDilemma(g Game) Game {
	//check if 2x2
	if g.Players.size > 2 {
		panic("game not supposed not have more than 2 players")
	}
	//load params

}
*/

// FindNash finds all NE present in the game
// ret value?
func (g *Game) FindNash(game Game) {

}

// FindParetto finds all NE present in the game
func (g *Game) FindParetto(game Game) {

}

//______________________________________________________________________________

// PrisonerDilemma subclass of game with specific parameter options
type PrisonerDilemma struct {
	Game
	tParam, rParam, sParam, pParam float64
}

// PDInterface interface for the prisoner dilemma
type PDInterface interface {
	CreatePDilemma([]*Action, []*Agent, float64, float64, float64, float64) *Game
}

// CreatePDilemma creates a pd game with the specified parameters
func (pd *PrisonerDilemma) CreatePDilemma(actions []*Action, players []*Agent, t, r, s, p float64) *Game {
	if len(players) > 2 {
		panic("must be 2 players only")
	}
	newGame := pd.CreateGame(actions, players)
	newGame.States[0][0].payoffs[0] = r
	newGame.States[0][0].payoffs[1] = r
	newGame.States[0][1].payoffs[0] = s
	newGame.States[0][1].payoffs[1] = t
	newGame.States[1][0].payoffs[0] = t
	newGame.States[1][0].payoffs[1] = s
	newGame.States[1][1].payoffs[0] = p
	newGame.States[1][1].payoffs[1] = p
	return newGame
}

//_______________________________EXTRAS________________________________________

// PrintGame generic function to print a 2x2 game
func PrintGame(g *Game) {

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print(i, j, ".....")
			fmt.Println(g.States[i][j].payoffs[0], g.States[i][j].payoffs[1])
		}
	}
}
