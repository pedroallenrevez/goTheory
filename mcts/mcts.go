package mcts

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// MCTS structure for the main algorithm - iterations, time etc.
type MCTS struct {
	C             float64 //exploration value?????????/
	CurrentState  WorldModel
	Root          *Node
	ExpandedNodes []*Node
}

// Interface Provides the interface for the MCTS algorithm. Run method and init
type Interface interface {
	Run(Node) Action
	Init(WorldModel)
}

//______________________________________________________________________________

// CreateMCTS - provides factory method for new MCTS algorithm
func CreateMCTS(state WorldModel) *MCTS {
	root := CreateNode(state, nil)
	new := MCTS{
		C:             1.4,
		CurrentState:  state,
		Root:          &root,
		ExpandedNodes: make([]*Node, 0),
	}
	return &new

}

// Run - returns the best action chosen by exploitation or exploration
func (mcts *MCTS) Run() *Action {
	// while within budget
	rand.Seed(time.Now().Unix())
	fmt.Println("--------------FRESH RUN----------------")
	for i := 0; i < 10; i++ {
		front := mcts.selection(mcts.Root)
		reward := mcts.playout(front.CurrentState)
		mcts.backPropagate(front, reward)
	}
	actions := mcts.exploitation(mcts.Root).CurrentState.Actions
	//reset expanded nodes
	return actions[len(actions)-1]
}

// Reset - resets the expanded nodes
func (mcts *MCTS) Reset() {
	mcts.ExpandedNodes = make([]*Node, 0)
}

func (mcts *MCTS) selection(initialNode *Node) *Node {
	node := initialNode
	for !(node.CurrentState.IsTerminal()) {
		if !(node.FullyExpanded()) {
			node = mcts.expand(node)
			break
		} else {
			node = mcts.exploration(node, mcts.C)
		}
	}
	return node
}

func (mcts *MCTS) expand(parent *Node) *Node {
	//generate child from parent
	//calculate new state from new executed action
	action := parent.UntriedActions[0]
	state := action.ApplyEffects(parent.CurrentState)
	//remove from untried actions
	parent.RemoveFrontAction()
	//add child to parent
	parent.AddChild(state) //add child creates node from state
	node := parent.ChildNodes[len(parent.ChildNodes)-1]
	mcts.ExpandedNodes = append(mcts.ExpandedNodes, node)
	return node

}

//there is state, currentstate, futurestate
func (mcts *MCTS) playout(state WorldModel) float64 {
	iterState := state
	for !iterState.IsTerminal() {
		//pick random action
		action := globalActions[rand.Intn(len(globalActions))]
		//applyeffects checks if its is executable gives the appropriate
		//reward and generates child
		state := action.ApplyEffects(iterState)
		//state = state with new actions
		iterState = state

	}
	return iterState.Fitness - state.Fitness
}

func (mcts *MCTS) backPropagate(node *Node, reward float64) {
	for node.Parent != nil {
		node.Visits++
		node.Reward += reward
		node = node.Parent
	}
	return
}

//______________________________________________________________________________

func (mcts *MCTS) exploration(node *Node, experimental float64) *Node {
	return mcts.bestChild(node, experimental)
}

func (mcts *MCTS) exploitation(node *Node) *Node {
	return mcts.bestChild(node, 0.0)

}

func (mcts *MCTS) bestChild(node *Node, experimental float64) *Node {
	bestScore := 0.0
	bestChildren := make([]*Node, 0)
	var bestChild *Node
	fmt.Print("n children: ")
	fmt.Println(len(node.ChildNodes))

	for _, child := range node.ChildNodes {
		exploitation := child.Reward / float64(child.Visits)
		exploration := math.Sqrt(math.Log(2*child.Visits) / float64(child.Visits))
		score := exploitation + exploration + experimental
		fmt.Println(score)
		child.PrintNode()
		if score == bestScore {
			bestChildren = append(bestChildren, child)
		}
		if score > bestScore {
			temp := make([]*Node, 0)
			bestChildren = append(temp, child)
			bestChild = child
			bestScore = score
		}
	}

	if len(bestChildren) == 0 {
		panic("no best children ???")
	}
	//return bestChildren[rand.Intn(len(bestChildren))]
	return bestChild

}

// PrintMcts - prints the mcts
func (mcts *MCTS) PrintMcts() {
	fmt.Println()
	fmt.Print("c: ")
	fmt.Println(mcts.C)
}

// UpdateExpandedMaps - updates
func (mcts *MCTS) UpdateExpandedMaps(w WorldMap) {
	for _, node := range mcts.ExpandedNodes {
		node.CurrentState.UpdateMap(*globalWorldMap)
	}
}
