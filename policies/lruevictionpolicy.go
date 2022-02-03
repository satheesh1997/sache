package policies

import "github.com/satheesh1997/sache/algorithms"

type (
	LRUEvictionPolicy struct {
		dll    algorithms.DoubleLinkedList
		mapper map[string]*algorithms.DoubleLinkedListNode
	}
)

func (lep *LRUEvictionPolicy) KeyAccessed(key string) {
	node, found := lep.mapper[key]

	if found {
		lep.dll.DetachNode(node)
		lep.dll.AddNodeAtLast(node)
	} else {
		lep.mapper[key] = lep.dll.AddElementAtLast(key)
	}
}

func (lep *LRUEvictionPolicy) EvictKey() string {
	first := lep.dll.First

	if first == nil {
		return ""
	}

	lep.dll.DetachNode(first)
	return first.Element
}

func NewLRUEvictionPolicy() LRUEvictionPolicy {
	policy := LRUEvictionPolicy{
		dll:    algorithms.DoubleLinkedList{First: nil, Last: nil},
		mapper: make(map[string]*algorithms.DoubleLinkedListNode),
	}
	return policy
}
