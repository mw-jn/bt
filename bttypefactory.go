package bt

var btNodeTypeFac btNodeTypeFactory

type objCreator func() IBTNode

type btNodeTypeFactory struct {
	objCreatorSet map[string]objCreator
}

func (ntf *btNodeTypeFactory) register(k string, newObj objCreator) {
	if btNodeTypeFac.objCreatorSet == nil {
		btNodeTypeFac.objCreatorSet = make(map[string]objCreator)
	}

	btNodeTypeFac.objCreatorSet[k] = newObj
}

func (ntf *btNodeTypeFactory) createBTNodeObject(k string) IBTNode {
	if newObj, ok := btNodeTypeFac.objCreatorSet[k]; ok && newObj != nil {
		return newObj()
	}
	return nil
}

func init() {
	btNodeTypeFac.register("selector", func() IBTNode { return &btSelector{} })
	btNodeTypeFac.register("selectorProbability", func() IBTNode { return &btSelectorProbability{} })
	btNodeTypeFac.register("selectorStochastic", func() IBTNode { return &btSelectorStochastic{} })
	btNodeTypeFac.register("sequence", func() IBTNode { return &btSequence{} })
	btNodeTypeFac.register("sequenceStochastic", func() IBTNode { return &btSequenceStochastic{} })
}
