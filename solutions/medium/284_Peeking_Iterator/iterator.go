package _84_Peeking_Iterator

type PeekingIterator struct {
	Iter      Iterator
	NextIndex int
}

func Constructor(iter Iterator) *PeekingIterator {
	return &PeekingIterator{
		Iter:      iter,
		NextIndex: -1,
	}
}

func (this *PeekingIterator) hasNext() bool {
	if this.NextIndex != -1 {
		return true
	}
	return this.Iter.hasNext()
}

func (this *PeekingIterator) next() int {
	if this.NextIndex != -1 {
		defer this.null()
		return this.NextIndex
	}
	return this.Iter.next()
}

func (this *PeekingIterator) peek() int {
	if this.NextIndex != -1 {
		return this.NextIndex
	}
	this.NextIndex = this.next()
	return this.NextIndex
}

func (this *PeekingIterator) null() {
	this.NextIndex = -1
}
