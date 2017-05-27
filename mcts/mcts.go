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
	C float64 //exploration value?????????/
	/*InProgress        bool
	MaxIterations     int
	MaxSelectionDepth int
	MaxPlayoutDepth   int
	*/
	BestFirstChild Node
	InitialNode    Node
	CurrentState   WorldModel
	//best action sequence?????
	//budget ??
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
func (mcts *MCTS) Init(state WorldModel) {

}

// Run - returns the best action chosen by exploitation or exploration
func (mcts *MCTS) Run(root Node) Action {
	// while within budget
	for i := 0; i < 100; i++ {
		front := mcts.selection(root)
		reward := mcts.playout(front.state)
		mcts.backPropagate(front, reward)
	}
	return mcts.bestChild(root)
}

func (mcts *MCTS) selection(initialNode Node) Node {
	var node Node
	for !initialNode.State.IsTerminal() {
		if !initialNode.FullyExpanded {
			initialNode = mcts.expand(initialNode)
			break
		} else {
			initialNode = mcts.bestUCTChild(initialNode)
		}
	}
	return initialNode
}

func (mcts *MCTS) expand(parent Node) {
	//calculate new state from new executed action
}

//there is state, currentstate, futurestate
func (mcts *MCTS) playout(state WorldModel) Reward {
	for !state.IsTerminal() {
		//pick random action
		//state = state with new actions
	}
}

func (mcts *MCTS) backPropagate(node Node, reward Reward) {
	for node != nil {
		node.Visits++
		node.Reward += reward
		node = node.Parent
	}
	return
}

//______________________________________________________________________________

func (mcts *MCTS) exploration(node Node, experimental float64) Node {
	return BestChild(node, experimental)
}

func (mcts *MCTS) exploitation(node Node) Node {
	return BestChild(node, 0.0)

}

func (mcts *MCTS) bestChild(node Node, experimental float64) Node {
	bestScore := 0
	bestChildren := make([]Node)
	for child := range node.ChildNodes {
		exploitation := child.Reward / child.Visits
		exploration := math.Sqrt(math.Log(2*child.Visits) / float64(child.Visits))
		score := exploitation + exploration + experimental
		if score == bestScore {
			bestChildren = append(bestChildren, child)
		}
		if score > bestScore {
			temp := make([]Node)
			bestChildren = append(temp, child)
			bestScore = score
		}
	}
	if len(bestChildren) == 0 {
		panic("no best children ???")
	}
	return bestChildren[rand.Intn(len(bestChildren))]

}
