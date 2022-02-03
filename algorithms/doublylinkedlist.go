package algorithms

type (
	DoubleLinkedListNode struct {
		Next    *DoubleLinkedListNode
		Prev    *DoubleLinkedListNode
		Element string
	}

	DoubleLinkedList struct {
		First *DoubleLinkedListNode
		Last  *DoubleLinkedListNode
	}
)

func (dll *DoubleLinkedList) DetachNode(node *DoubleLinkedListNode) {
	if node == dll.First {
		first := dll.First.Next
		if first != nil {
			first.Prev = nil
		}
	} else if node == dll.Last {
		last := dll.Last.Prev
		if last != nil {
			last.Next = nil
		}
	} else {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
	}
}

func (dll *DoubleLinkedList) AddNodeAtLast(node *DoubleLinkedListNode) {
	if node == nil {
		panic("Node is nil, Cannot add it to the list.")
	}

	if dll.Last == nil {
		dll.Last = node
		dll.First = node
	} else {
		dll.Last.Next = node
		node.Prev = dll.Last
		node.Next = nil
		dll.Last = node
	}
}

func (dll *DoubleLinkedList) AddElementAtLast(element string) *DoubleLinkedListNode {
	node := &DoubleLinkedListNode{Element: element}
	dll.AddNodeAtLast(node)
	return node
}
