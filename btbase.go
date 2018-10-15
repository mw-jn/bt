package bt

// NodeType can be redefined type.
type NodeType int

// NodeStatusType can be redefined type.
type NodeStatusType int

const (
	// NodeTypeSequence show type sequence
	NodeTypeSequence NodeType = iota
	// NodeTypeSelector show type selector
	NodeTypeSelector
	// NodeTypeParallel show type parallel
	NodeTypeParallel
	// NodeTypeCondition show type condition
	NodeTypeCondition
	// NodeTypeAction show type Action
	NodeTypeAction
	// NodeTypeCustom show type custom
	NodeTypeCustom
)

const (
	// NodeStatusTypeRunning defalut state
	NodeStatusTypeRunning NodeStatusType = iota
	// NodeStatusTypeSuccess show success state
	NodeStatusTypeSuccess
	// NodeStatusTypeFailure show failure state
	NodeStatusTypeFailure
)

var _ IBTNode = &BTBase{}

// IBTNode supply a interface for BTNode in behavior tree.
type IBTNode interface {
	InitNode(name string, data interface{})
	Tick() NodeStatusType // 更新信息
	Name() string         // 节点名
	Type() NodeType       // node type
	AddChild(i IBTNode)
	SetParent(i IBTNode)                  // set parent node
	ForeachNode(f func(i IBTNode))        //
	ForeachNodeIf(f func(i IBTNode) bool) //
}

// BTBase implements the base data structure for behavior tree.
type BTBase struct {
	nodeName string
	category NodeType
	parent   IBTNode
	children []IBTNode
}

// NewBTBaseNode return BTBase structure.
func NewBTBaseNode(t NodeType) *BTBase {
	b := &BTBase{
		category: t,
	}
	return b
}

// InitNode implement IBTNode method.
func (b *BTBase) InitNode(name string, data interface{}) {
	b.nodeName = name
}

// Type return the node type.
func (b *BTBase) Type() NodeType {
	return b.category
}

// Name implement IBNode method.
func (b *BTBase) Name() string {
	return b.nodeName
}

// Tick update
func (b *BTBase) Tick() NodeStatusType {
	return NodeStatusTypeRunning
}

// ChildrenCount return the size of children node.
func (b *BTBase) ChildrenCount() uint {
	return uint(len(b.children))
}

// AddChild add i to children list.
func (b *BTBase) AddChild(i IBTNode) {
	b.children = append(b.children, i)
	i.SetParent(b)
}

// SetParent set node i as parent.
func (b *BTBase) SetParent(i IBTNode) {
	b.parent = i
}

// ForeachNode traverse all nodes.
func (b *BTBase) ForeachNode(f func(i IBTNode)) {
	for _, node := range b.children {
		f(node)
	}
}

// ForeachNodeIf traverse all nodes and stop tranvese if find the need one.
func (b *BTBase) ForeachNodeIf(f func(i IBTNode) bool) {
	for _, node := range b.children {
		if f(node) {
			return
		}
	}
}
