package main

import (
	"fmt"
	"github.com/pedroallenrevez/goTheory/gameTheory"
)

func main() {
	var impl = new(gameTheory.PrisonerDilemma)
	var g gameTheory.PDInterface = impl
	//create actions
	var actionJackson = new(gameTheory.Action)
	var a gameTheory.AInterface = actionJackson
	action1 := a.CreateAction("yolo")
	action2 := a.CreateAction("swag")
	actArray := make([]*gameTheory.Action, 2)
	actArray[0] = action1
	actArray[1] = action2
	//create players
	var agentJackson = new(gameTheory.Agent)
	var ag gameTheory.AgInterface = agentJackson
	var strat gameTheory.Strategy = func() *gameTheory.Action {
		return action1
	}
	p1 := ag.CreateAgent(actArray, strat)
	strat = func() *gameTheory.Action {
		return action2
	}
	p2 := ag.CreateAgent(actArray, strat)
	players := make([]*gameTheory.Agent, 2)
	players[0] = p1
	players[1] = p2

	pd := g.CreatePDilemma(actArray, players, 1.8, 1.0, 0.0, 0.0)
	gameTheory.PrintGame(pd)
}
