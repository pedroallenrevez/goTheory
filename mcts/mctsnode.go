package mcts

import (
	"fmt"
	"github.com/pedroallenrevez/goTheory/gameTheory/action"
)

// Node node structure for the MCTS. holds the state
type Node struct {
	CurrentState State
	Parent       Node
	Action       action.Action // ? is it needed
	ChildNodes   []Node
	Visits       float64
	Reward       float64
	H            float64
	AccumulatedE float64
}

// CreateNode - initialize a node with a state, and a parent
func (n *Node) CreateNode(state State, parent Node) *Node {
	nodeList := make([]Node)
	new := Node{
		CurrentState: state,
		Parent:       parent,
		Visits:       0,
		ChildNodes:   nodeList,
		Reward:       0,
		H:            0,
		AccumulatedE: 0,
	}
	return &new

}

// AddChild - Create a child node from state, under the current node
func (n *Node) AddChild(state State) {
	child := CreateNode(state, n)
	n.ChildNodes = append(n.ChildNodes, child)
}

// Update - Update the Node with the reward and number of visits
func (n *Node) Update(r Reward) {
	n.Reward += r
	n.Visits++
}

// FullyExpanded - if the number of children is equal to the n actions, its
// fully expanded
func (n Node) FullyExpanded() bool {
	if len(n.ChildNodes) == len(n.State.Game.Actions) {
		return true
	}
	return false

}
