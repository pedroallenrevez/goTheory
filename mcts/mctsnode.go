package mcts

import (
	"fmt"
)

// Node node structure for the MCTS. holds the state
type Node struct {
	CurrentState   WorldModel
	Parent         *Node
	ChildNodes     []*Node
	UntriedActions []*Action
	Visits         float64
	Reward         float64
}

// CreateNode - initialize a node with a state, and a parent
func CreateNode(state WorldModel, parent *Node) Node {
	nodeList := make([]*Node, 0)
	//TODO assign actions to the Node
	new := Node{
		CurrentState:   state,
		Parent:         parent,
		Visits:         0,
		ChildNodes:     nodeList,
		Reward:         0,
		UntriedActions: globalActions,
	}
	return new

}

// AddChild - Create a child node from state, under the current node
func (n *Node) AddChild(state WorldModel) {
	fmt.Println("adding child")
	nodeList := make([]*Node, 0)
	child := &Node{
		CurrentState:   state,
		Parent:         n,
		Visits:         0,
		ChildNodes:     nodeList,
		Reward:         0,
		UntriedActions: globalActions,
	}
	n.ChildNodes = append(n.ChildNodes, child)
}

// RemoveFrontAction - removes the first action of untried actions. used for
// expanding nodes
func (n *Node) RemoveFrontAction() {
	n.UntriedActions = append(n.UntriedActions[:0], n.UntriedActions[1:]...)

}

// Update - Update the Node with the reward and number of visits
func (n *Node) Update(r float64) {
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

// PrintNode - prints the node
func (n *Node) PrintNode() {
	fmt.Println("state: ")
	n.CurrentState.PrintState()
	fmt.Print("parent: ")
	if n.Parent == nil {
		fmt.Print("nil")
	} else {
		fmt.Print("not nil")
	}
	fmt.Print("childnodes: ")
	fmt.Print(len(n.ChildNodes))
	fmt.Print("untried actions: ")
	fmt.Print(len(n.UntriedActions))
	fmt.Print("visists: ")
	fmt.Print(n.Visits)
	fmt.Print("reward: ")
	fmt.Print(n.Reward)
	fmt.Println()
}
