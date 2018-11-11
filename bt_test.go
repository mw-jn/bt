package bt

import (
	"fmt"
	"os"
	"testing"
)

func TestBTCondition(t *testing.T) {
	rootPath := os.Getenv("GOPATH")
	var behaviortree behaviorTree
	err := behaviortree.loadBehaviorTree(rootPath + "/src/bt/test_data/sample.xml")
	//err := behaviortree.loadBehaviorTree("../test_data/sample.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	behaviortree.Print(behaviortree.root)
}
