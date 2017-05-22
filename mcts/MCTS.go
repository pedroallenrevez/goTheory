package main

import "fmt"

type reward struct {
	value float64
}

func (r *reward) RewardFromNode(node MCTSNode) float64 {

}

type MCTSInterface interface {
	Run() WorldModel.Action
	Init(WorldModel.CurrentState)
}

type MCTSNode struct {
}

type MCTS struct {
}

func (mcts *MCTS) Init(state WorldModel.CurrentState) {

}

func (mcts *MCTS) Run() WorldModel.Action {

}

func (mcts *MCTS) selection(initialNode MCTSNode) MCTSNode {

}

func (mcts *MCTS) expand(parent MCTSNode, action WorldModel.Action) {

}

//there is state, currentstate, futurestate
func (mcts *MCTS) playout(initialPlayout WorldModel.State) Reward {

}

func (mcts *MCTS) backPropagate(node MCTSNode, reward Reward) {

}

//______________________________________________________________________________

func (mcts *MCTS) bestUCTChild(node MCTSNode) MCTSNode {

}

func (mcts *MCTS) bestChild(node MCTSNode) MCTSNode {

}
