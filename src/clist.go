package main

type Object interface{}

type ObjectPtr interface{}

// an function type for match data
type MatchFun func(data1 Object, data2 Object) int

func whatType(x interface{}) string {
	switch x.(type) {
	case int:
		return "int"
	case string:
		return "字符串"
	default:
		return "未知"
	}
}

type CNode struct {
	data Object
	next *CNode
}

type CList struct {
	size uint64
	head *CNode
}

func (cList *CList) Init() {
	lst := *cList
	lst.size = 0
	lst.head = nil
}

func (cList *CList) Append(data Object) bool {
	node := new(CNode)
	(*node).data = data

	if cList.GetSize() == 0 { // empty
		(*cList).head = node
	} else {
		item := cList.GetHead()
		for ; (*item).next != cList.GetHead(); item = (*item).next {
		}
		(*item).next = node
	}

	(*node).next = (*cList).head
	(*cList).size++

	return true
}

func (cList *CList) InsertNext(elmt *CNode, data Object) bool {
	if elmt == nil {
		return false
	}

	node := new(CNode)
	(*node).data = data

	(*node).next = (*elmt).next
	(*elmt).next = node

	(cList).size++

	return true
}

func (cList *CList) Remove(elmt *CNode) Object {
	if elmt == nil {
		return false
	}

	item := cList.GetHead()
	for ; (*item).next != elmt; item = (*item).next {
	}

	(*item).next = (*elmt).next
	(*cList).size--

	return elmt.GetData()
}

func (cList *CList) GetSize() uint64 {
	return (*cList).size
}

func (cList *CList) GetHead() *CNode {
	return (*cList).head
}

func (node *CNode) GetData() Object {
	return (*node).data
}

func (node *CNode) GetNext() *CNode {
	return (*node).next
}
