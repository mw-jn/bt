package bt

import (
	"fmt"
	"xxml"
)

type behaviorTree struct {
	root IBTNode // root

}

func (bht *behaviorTree) loadBehaviorTree(filepath string) error {
	xmlDoc, err := xxml.NewXMLDocByFilePath(filepath)
	if err != nil {
		return err
	}
	method := func(p interface{}, i xxml.IXMLNode) interface{} {
		parent, ok := p.(IBTNode)
		if !ok {
			panic(fmt.Sprintf("Copy Tree error: (type error, %T)", p))
		}
		obj := btNodeTypeFac.createBTNodeObject(i.Name())
		if obj == nil {
			return nil
		}
		obj.loadAttributes(i.Name(), i)
		parent.AddChild(obj)
		return obj
	}
	bht.root = &btNodeBase{}
	bht.root.loadAttributes("Root", nil)
	xmlDoc.CopyTree(bht.root, method)
	return nil
}

func (bht *behaviorTree) Print(root IBTNode) {
	fmt.Printf("Node Name:%v\n", root.Name())
	fmt.Printf("Node Attributs:%v\n", root.printAttributes())
	root.ForeachNode(func(i IBTNode) {
		bht.Print(i)
	})
}
