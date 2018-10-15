package bt

var _ IBTNode = &BTSequence{}

// BTSequence respresents sequence node.
type BTSequence struct {
	*BTBase
}

// InitNode implement the interface IBTNode.
func (bts *BTSequence) InitNode(name string, data interface{}) {
	// init common data
	bts.BTBase = NewBTBaseNode(NodeTypeSequence)
	bts.BTBase.InitNode(name, data)
	// init BTSequence self data
	// ...
}

// Tick implement the interface IBTNode.
func (bts *BTSequence) Tick() NodeStatusType {
	s := NodeStatusTypeSuccess
	bts.ForeachNodeIf(func(i IBTNode) bool {
		childStatus := i.Tick()
		if childStatus == NodeStatusTypeSuccess {
			return false
		}

		s = childStatus
		return true
	})
	return s
}
