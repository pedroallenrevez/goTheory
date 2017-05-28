package mcts

// WorldModel - worldmodel abstraction for sk8rboy
type WorldModel struct {
	//score
	SID       int
	Map       []Cell
	Fitness   float64
	TurnsLeft int
}

// IsTerminal - checks if node is terminal or not
func (w *WorldModel) IsTerminal() bool {
	if w.TurnsLeft == 0 {
		return true
	}
	return false

}

// GenerateChild - gens a child world model
func (w *WorldModel) GenerateChild() *WorldModel {
	newState := &WorldModel{
		SID:       w.SID,
		Fitness:   w.Fitness + 1, //fitness + action reward
		TurnsLeft: w.TurnsLeft - 1,
	}
	return newState
}

// Reward - Gets Reward for model
func (w *WorldModel) Reward() Reward {
	//TODO here should be checked if move was done or not
	// if exec r1 if not r2
	var newReward Reward
	newReward = 0
	return newReward
}
