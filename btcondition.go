package bt

// BTConditionBase respresents condition node.
type BTConditionBase struct {
	*BTBase
}

func (btc *BTConditionBase) initData(name string, data interface{}) {
	btc.BTBase = NewBTBaseNode(NodeTypeSelector)
	btc.InitNode(name, data)
}
