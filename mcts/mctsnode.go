package mcts

import (
//"fmt"
)

// Node node structure for the MCTS. holds the state
type Node struct {
	CurrentState   *WorldModel
	Parent         *Node
	Action         *Action // action that brought to this state
	ChildNodes     []*Node
	UntriedActions []*Action
	Visits         float64
	Reward         float64
	// for biased mcts
	H            float64
	AccumulatedE float64
}

// CreateNode - initialize a node with a state, and a parent
func (n *Node) CreateNode(state *WorldModel, parent *Node) *Node {
	nodeList := make([]*Node, 0)
	//TODO assign actions to the Node
	//TODO assign action that led to this state
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
func (n *Node) AddChild(state *WorldModel) {
	child := n.CreateNode(state, n)
	n.ChildNodes = append(n.ChildNodes, child)
}

// RemoveFrontAction - removes the first action of untried actions. used for
// expanding nodes
func (n *Node) RemoveFrontAction() {
	n.UntriedActions = append(n.UntriedActions[:0], n.UntriedActions[1:]...)

}

// Update - Update the Node with the reward and number of visits
func (n *Node) Update(r Reward) {
	n.Reward += float64(r)
	n.Visits++
}

// FullyExpanded - if the number of children is equal to the n actions, its
// fully expanded
func (n Node) FullyExpanded() bool {
	if len(n.UntriedActions) == 0 {
		return true
	}
	return false

}
