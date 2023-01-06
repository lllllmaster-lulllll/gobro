package list

import "fmt"

type LinkNode struct {
	Data int64
	Next *LinkNode
}

func AddLinkNode() {
	// Create a new LinkNode
	node := &LinkNode{
		Data: 2,
	}

	node1 := new(LinkNode)
	node1.Data = 3
	node.Next = node1

	node2 := new(LinkNode)
	node2.Data = 4
	node1.Next = node2

	// todo 按顺序打印数据
	currentNode := node
	for {
		if currentNode != nil {
			fmt.Println(currentNode.Data)

			currentNode = currentNode.Next
			continue
		}
		break

	}

}
