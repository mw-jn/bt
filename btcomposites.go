package bt

import (
	"fmt"
	"math/rand"
)

const (
	invalidChildIndex = -1
)

// IBTCompositeNode add new methods
type IBTCompositeNode interface {
	IBTNode
	onEnter()
	onLeave()
}

type btCompositeBase struct {
	btNodeBase
	activeChildIndex int // 当前索
}

func (cb *btCompositeBase) onEnter() {

}

func (cb *btCompositeBase) onLeave() {

}

func (cb *btCompositeBase) checkActiveIndexInvalid() bool {
	return cb.activeChildIndex != invalidChildIndex
}

func (cb *btCompositeBase) activeChildNode() IBTNode {
	return cb.childByIndex(cb.activeChildIndex)
}

// Selector
// At least one child
type btSelector struct {
	btCompositeBase
}

func (slc *btSelector) onEnter() {
	if slc.childrenCount() <= 0 {
		panic("BT Node Selector error: No Child.")
	}
	slc.activeChildIndex = 0
}

func (slc *btSelector) Tick(childStatus NodeStatusType) NodeStatusType {
	if slc.activeChildIndex >= slc.childrenCount() {
		panic("BT Node Selector error: Child Index Overflow.")
	}

	s := childStatus

	for {
		if s == NodeStatusTypeRunning {
			childNode := slc.activeChildNode()
			s = slc.exec(childNode)
		}

		if s != NodeStatusTypeFailure {
			return s
		}

		slc.activeChildIndex++
		if slc.activeChildIndex >= slc.childrenCount() {
			return NodeStatusTypeFailure
		}

		s = NodeStatusTypeRunning
	}
} // end Selector

type btSelectorProbability struct {
	btCompositeBase

	weight      []float32
	totalWeight float32
}

func (sp *btSelectorProbability) onEnter() {
	sp.activeChildIndex = invalidChildIndex

	sp.weight = sp.weight[:0]
	sp.totalWeight = 0

	for _, node := range sp.children {
		if i, ok := node.(interface{}); ok { // TODO
			w := i.(float32) // TODO
			sp.weight = append(sp.weight, w)
			sp.totalWeight += w
		} else {
			panic(fmt.Sprintf("BT Node SelectorProbability error: Child Node Type error, %T", node))
		}
	}
}

func (sp *btSelectorProbability) onExit() {
	sp.activeChildIndex = invalidChildIndex
}

func (sp *btSelectorProbability) Tick(childStatus NodeStatusType) NodeStatusType {
	s := childStatus
	if s != NodeStatusTypeRunning {
		return s
	}

	if sp.checkActiveIndexInvalid() {
		childNode := sp.activeChildNode()
		s = sp.exec(childNode)
		return s
	}

	// 否则产生随机数
	r := sp.totalWeight * rand.Float32()

	sum := float32(0)

	for i := 0; i < sp.childrenCount(); i++ {
		w := sp.weight[i]
		sum += w
		if w > 0 && sum >= r {
			childNode := sp.childByIndex(i)
			s = sp.exec(childNode)
			if s == NodeStatusTypeRunning {
				sp.activeChildIndex = i
			}
			return s
		}

	}

	return NodeStatusTypeFailure
} // end SelectorProbability

type btSelectorStochastic struct {
	btCompositeBase
}

// Sequence
type btSequence struct {
	btCompositeBase
}

func (sq *btSequence) onEnter() {
	if sq.childrenCount() <= 0 {
		panic("BT Node Sequence error: No Child.")
	}
	sq.activeChildIndex = 0
}

func (sq *btSequence) Tick(childStatus NodeStatusType) NodeStatusType {
	if sq.activeChildIndex >= sq.childrenCount() {
		panic("BT Node Sequence error: Child Index Overflow.")
	}

	s := childStatus
	for {
		if s == NodeStatusTypeRunning {
			childNode := sq.activeChildNode()
			s = sq.exec(childNode)
		}

		if s != NodeStatusTypeSuccess {
			return s
		}

		sq.activeChildIndex++
		if sq.activeChildIndex >= sq.childrenCount() {
			return NodeStatusTypeSuccess
		}

		s = NodeStatusTypeRunning
	}
} // end Sequence
