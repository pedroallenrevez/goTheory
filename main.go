package main

import (
	"fmt"
	/*"github.com/pedroallenrevez/goTheory/gameTheory"*/
	"github.com/pedroallenrevez/goTheory/mcts"
)

func main() {
	/*
		var impl = new(gameTheory.PrisonerDilemma)
		var g gameTheory.PDInterface = impl
		//create actions
		var actionJackson = new(action.Action)
		var a action.AInterface = actionJackson
		action1 := a.CreateAction("yolo")
		action2 := a.CreateAction("swag")
		actArray := make([]*action.Action, 2)
		actArray[0] = action1
		actArray[1] = action2
		//create players
		var agentJackson = new(agent.Agent)
		var ag agent.AgInterface = agentJackson
		var strat agent.Strategy = func() *action.Action {
			return action1
		}
		p1 := ag.CreateAgent(actArray, strat)
		strat = func() *action.Action {
			return action2
		}
		p2 := ag.CreateAgent(actArray, strat)
		players := make([]*agent.Agent, 2)
		players[0] = p1
		players[1] = p2
		//create game
		pd := g.CreatePDilemma(actArray, players, 1.8, 1.0, 0.0, 0.0)
		gameTheory.PrintGame(pd)
		pd.Solve()
		fmt.Println(p1.TotalScore)
		fmt.Println(p2.TotalScore)
		//g.LoadPrisonerDilemma(pd, 1.9, 1.0, 0.0, 0.0)
	*/
	var impl = new(mcts.GameManager)
	var gman mcts.GManInterface = impl

	// width height nskaters speed radius r1 r2
	gman.Init(10, 10, 5, 1, 1, 1, 0)

	NumTurns := 100
	for x := 0; x < NumTurns; x++ {
		gman.PrintGame()
		gman.PrintSkaters()
		gman.Update()
		fmt.Println(x)
	}
}
