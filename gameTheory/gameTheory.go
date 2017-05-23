package gameTheory

import (
	"fmt"
	"github.com/pedroallenrevez/goTheory/gameTheory/action"
	"github.com/pedroallenrevez/goTheory/gameTheory/agent"
)

//__________________________________CORE________________________________________

// Game structure for a normal n x n games
type Game struct {
	Actions []*action.Action
	Players []*agent.Agent
	//only has a dictionary of dictionary of states
	//TODO implement nmap with n players
	//2 player implementation
	States [][]state
}

// GInterface interfface responsible for the game method available
type GInterface interface {
	CreateGame([]action.Action, []agent.Agent) *Game
	Solve(Game)
}

type state struct {
	//payoffs are ordered according to player ID
	payoffs [2]float64
}

// CreateGame creates a game with n action.Actions, m players
func (g *Game) CreateGame(actions []*action.Action, players []*agent.Agent) *Game {
	// maps of maps dictionaries
	// 2x2 prisoners dilemma
	// 1x1 or NxN

	//init states with size of action.Actions we want a.len x a.len

	NIDstates := make([][]state, len(actions))

	for i := range NIDstates {
		NIDstates[i] = make([]state, len(actions))
	}

	//payoffs := make([2]float64, len(players))
	var payoffs [2]float64

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

// CalculatePayoff returns payoff for given game players and action.Action
// it is not needed to compute all the values at the beggining, because the
// choice of action.Action is deterministic. We just calculate it when needed.
func (g *Game) CalculatePayoff(game *Game, p *agent.Agent, act *action.Action) {
	//run q function
}

// Solve returns payoff for each agent.Agent in game
func (g *Game) Solve() {
	//add score to each participating agent.Agent
	//for all agent.Agents - calculate payoff

	//actual implementation will be much different
	payoff := g.States[0][1].payoffs
	g.Players[0].TotalScore += payoff[0]
	g.Players[1].TotalScore += payoff[1]

	//DO STUFF
}

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
	CreatePDilemma([]*action.Action, []*agent.Agent, float64, float64, float64, float64) *Game
	//LoadPrisonerDilemma(*Game, float64, float64, float64, float64)
}

// CreatePDilemma creates a pd game with the specified parameters
func (pd *PrisonerDilemma) CreatePDilemma(actions []*action.Action, players []*agent.Agent, t, r, s, p float64) *Game {
	if len(players) > 2 {
		panic("must be 2 players only")
	}
	newGame := pd.CreateGame(actions, players)
	var arr [2]float64
	arr[0] = r
	arr[1] = r
	newGame.States[0][0].payoffs = arr
	arr[0] = s
	arr[1] = t
	newGame.States[0][1].payoffs = arr
	arr[0] = t
	arr[1] = s
	newGame.States[1][0].payoffs = arr
	arr[0] = p
	arr[1] = p
	newGame.States[1][1].payoffs = arr
	return newGame
}

/*
// LoadPrisonerDilemma receive a game and assign the PD values to them
// THIS ONLY WILL WORK FOR 2x2 GAMES
func (pd *PrisonerDilemma) LoadPrisonerDilemma(game *Game, t, r, p, s float64) {
	//check if 2x2
	if len(game.Players) > 2 {
		panic("game not supposed not have more than 2 players")
	}
	//load params
	pd.States[0][0].payoffs[0] = r
	pd.States[0][0].payoffs[1] = r
	pd.States[0][1].payoffs[0] = s
	pd.States[0][1].payoffs[1] = t
	pd.States[1][0].payoffs[0] = t
	pd.States[1][0].payoffs[1] = s
	pd.States[1][1].payoffs[0] = p
	pd.States[1][1].payoffs[1] = p

}
*/

//_______________________________EXTRAS________________________________________

// PrintGame generic function to print a 2x2 game
func PrintGame(g *Game) {

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print(i, j, "...")
			fmt.Println(g.States[i][j].payoffs[0], g.States[i][j].payoffs[1])
		}
	}
}
