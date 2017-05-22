package main

import (
	"fmt"
	"string"
)

type State struct {
	actions []Action
	parent  State
	game    GT.Game
}
type CurrentState struct {
}

type Action struct {
	name string
}
