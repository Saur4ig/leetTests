package _41_Flatten_Nested_List_Iterator

type NestedIterator struct {
	sorted       []*NestedInteger
	nextPosition int
	len          int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	nI := make([]*NestedInteger, 0)

	for _, nestedInt := range nestedList {
		if !nestedInt.IsInteger() {
			nI = append(nI, fromList(nestedInt)...)
		} else {
			nI = append(nI, nestedInt)
		}
	}
	return &NestedIterator{
		sorted:       nI,
		nextPosition: 0,
		len:          len(nI),
	}
}

func fromList(n *NestedInteger) []*NestedInteger {
	nI := make([]*NestedInteger, 0)
	for _, val := range n.GetList() {
		if !val.IsInteger() {
			nI = append(nI, fromList(val)...)
		} else {
			nI = append(nI, val)
		}
	}
	return nI
}

func (this *NestedIterator) Next() int {
	defer this.increase()
	return this.sorted[this.nextPosition].GetInteger()
}

func (this *NestedIterator) increase() {
	this.nextPosition++
}

func (this *NestedIterator) HasNext() bool {
	if this.nextPosition < this.len {
		return true
	}
	return false
}
