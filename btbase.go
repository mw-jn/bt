package bt

import (
	"fmt"
	"xxml"
)

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
	// NodeStatusTypeInvalid show invalid
	NodeStatusTypeInvalid NodeStatusType = iota
	// NodeStatusTypeSuccess show success state
	NodeStatusTypeSuccess
	// NodeStatusTypeFailure show failure state
	NodeStatusTypeFailure
	// NodeStatusTypeRunning defalut state
	NodeStatusTypeRunning
)

var _ IBTNode = &btNodeBase{}

// IBTNode supply a interface for BTNode in behavior tree.
type IBTNode interface {
	loadAttributes(name string, data interface{})
	Tick(agent *Agent, childStatus NodeStatusType) NodeStatusType // 更新信息
	Name() string                                                 // 节点名
	//	Type() NodeType                                               // node type
	AddChild(i IBTNode)
	SetParent(i IBTNode)           // set parent node
	ForeachNode(f func(i IBTNode)) //
	//	ForeachNodeIf(f func(i IBTNode) bool) //
	// Clone() IBTNode
	printAttributes() string
}

// btNodeBase implements the base data structure for behavior tree.
type btNodeBase struct {
	nodeName string
	//	category NodeType

	parent   IBTNode
	children []IBTNode

	attr attributes // 节点属性
}

func (b *btNodeBase) loadAttributes(name string, data interface{}) {
	b.nodeName = name

	if data == nil {
		return
	}

	d, ok := data.(xxml.IXMLNode)
	if !ok {
		panic(fmt.Sprintf("BT Node Base init error: (original data type %T) => (*xxml.XMLNode)", data))
	}

	d.ForeachAttr(func(k, v string) {
		b.attr.add(k, v)
	})
}

/*
// Type return the node type.
func (b *btNodeBase) Type() NodeType {
	return b.category
}
*/

// Name implement IBNode method.
func (b *btNodeBase) Name() string {
	return b.nodeName
}

// Tick update
func (b *btNodeBase) Tick(agent *Agent, childStatus NodeStatusType) NodeStatusType {
	return NodeStatusTypeRunning
}

// childrenCount return the size of children node.
func (b *btNodeBase) childrenCount() int {
	return len(b.children)
}

// AddChild add i to children list.
func (b *btNodeBase) AddChild(i IBTNode) {
	b.children = append(b.children, i)
	i.SetParent(b)
}

func (b *btNodeBase) childByIndex(i int) IBTNode {
	return b.children[i]
}

// SetParent set node i as parent.
func (b *btNodeBase) SetParent(i IBTNode) {
	b.parent = i
}

// 转交给子节点执行
func (b *btNodeBase) dispatchExec(i IBTNode, agent *Agent) NodeStatusType {
	childStatus := NodeStatusTypeRunning
	return i.Tick(agent, childStatus)
}

// ForeachNode traverse all nodes.
func (b *btNodeBase) ForeachNode(f func(i IBTNode)) {
	for _, node := range b.children {
		f(node)
	}
}

/*
// ForeachNodeIf traverse all nodes and stop tranvese if find the need one.
func (b *btNodeBase) ForeachNodeIf(f func(i IBTNode) bool) {
	for _, node := range b.children {
		if f(node) {
			return
		}
	}
}
*/

func (b *btNodeBase) printAttributes() string {
	return b.attr.String()
}
