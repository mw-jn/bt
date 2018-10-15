package bt

var _ IBTNode = &BTSelector{}

// BTSelector respresents selector node.
type BTSelector struct {
	*BTBase
}

// InitNode implement the interface IBTNode.
func (btst *BTSelector) InitNode(name string, data interface{}) {
	// init common data
	btst.BTBase = NewBTBaseNode(NodeTypeSelector)
	btst.BTBase.InitNode(name, data)
}

// Tick implement the interface IBTNode.
func (btst *BTSelector) Tick() NodeStatusType {
	s := NodeStatusTypeFailure
	btst.ForeachNodeIf(func(i IBTNode) bool {
		childState := i.Tick()
		if childState == NodeStatusTypeFailure {
			return false
		}
		s = childState
		return true
	})
	return s
}
