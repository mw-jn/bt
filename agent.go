package bt

// IBTAgent define a interface for behavior tree.
type IBTAgent interface {
}

// Agent implements proxy for behavior tree.
type Agent struct {
	IBTAgent
}

// NewAgent return agent fo behavior tree.
func NewAgent(a IBTAgent) *Agent {
	return &Agent{a}
}
