package mcts

import "fmt"

// Reward - specification for a reward
type Reward float64

// RewardFromNode gets reward from the node - maybe not needed in this context
// yet
func (r *reward) RewardFromNode(node Node) reward {
	return 0.0
}
