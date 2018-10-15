package bt

const (
	defaultM = 2
)

// BTParallel respresents parallel node.
type BTParallel struct {
	*BTBase
}

func (btp *BTParallel) initData(name string, data interface{}) {
	// init common data
	btp.BTBase = NewBTBaseNode(NodeTypeParallel)
	btp.BTBase.InitNode(name, data)
}

// Tick implement the interface IBTNode.
func (btp *BTParallel) Tick() NodeStatusType {
	succCnt := uint(0)
	failCnt := uint(0)
	btp.ForeachNode(func(i IBTNode) {
		switch i.Tick() {
		case NodeStatusTypeSuccess:
			succCnt++
		case NodeStatusTypeFailure:
			failCnt++
		}
	})

	if succCnt >= defaultM {
		return NodeStatusTypeSuccess
	} else if failCnt+defaultM > btp.ChildrenCount() {
		return NodeStatusTypeFailure
	}
	return NodeStatusTypeSuccess
}
