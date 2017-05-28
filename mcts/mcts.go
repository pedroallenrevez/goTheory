package mcts

import (
	//"fmt"
	"math"
	"math/rand"
	"time"
)

// Reward - abstraction for rewards
type Reward float64

// MCTS structure for the main algorithm - iterations, time etc.
type MCTS struct {
	C            float64 //exploration value?????????/
	CurrentState WorldModel
	Actions      []*Action
}

// Interface Provides the interface for the MCTS algorithm. Run method and init
type Interface interface {
	Run(Node) Action
	Init(WorldModel)
	CreateMCTS() *MCTS
}

//______________________________________________________________________________

// CreateMCTS - provides factory method for new MCTS algorithm
func CreateMCTS(state WorldModel) *MCTS {
	rand.Seed(time.Now().Unix())
	new := MCTS{
		C: 1.4,
	}
	return &new

}

// Init - initialize the MCTS with an initial world state
func (mcts *MCTS) Init(state WorldModel, actions []*Action) {
	mcts.CurrentState = state
	mcts.C = 1.4
	mcts.Actions = actions
}

// Run - returns the best action chosen by exploitation or exploration
func (mcts *MCTS) Run(root *Node) *Action {
	// while within budget
	for i := 0; i < 100; i++ {
		front := mcts.selection(root)
		reward := mcts.playout(front.CurrentState)
		mcts.backPropagate(front, reward)
	}
	return mcts.exploitation(root).Action
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
	newChild := parent.CurrentState.GenerateChild()
	//calculate new state from new executed action
	parent.UntriedActions[0].ApplyEffects(newChild)
	//remove from untried actions
	parent.RemoveFrontAction()
	//add child to parent
	parent.AddChild(newChild) //add child already created node from parent
	return parent.ChildNodes[len(parent.ChildNodes)-1]

}

//there is state, currentstate, futurestate
func (mcts *MCTS) playout(state *WorldModel) Reward {
	for !state.IsTerminal() {
		child := state.GenerateChild()
		//pick random action
		action := CreateAction("ahah")
		action.ApplyEffects(child)
		//state = state with new actions
		state = child

	}
	return state.Reward()
}

func (mcts *MCTS) backPropagate(node *Node, reward Reward) {
	for node != nil {
		node.Visits++
		node.Reward += float64(reward)
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
	for _, child := range node.ChildNodes {
		exploitation := child.Reward / float64(child.Visits)
		exploration := math.Sqrt(math.Log(2*child.Visits) / float64(child.Visits))
		score := exploitation + exploration + experimental
		if score == bestScore {
			bestChildren = append(bestChildren, child)
		}
		if score > bestScore {
			temp := make([]*Node, 0)
			bestChildren = append(temp, child)
			bestScore = score
		}
	}
	if len(bestChildren) == 0 {
		panic("no best children ???")
	}
	return bestChildren[rand.Intn(len(bestChildren))]

}
