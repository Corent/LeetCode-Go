package algorithm284

type Iterator struct{}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
	return false
}

func (this *Iterator) next() int {
	// Returns the next element in the iteration.
	return 0
}

type PeekingIterator struct {
	iter  *Iterator
	queue []int
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iter: iter, queue: make([]int, 0)}
}

func (this *PeekingIterator) hasNext() bool {
	if len(this.queue) == 0 {
		return this.iter.hasNext()
	}
	return true
}

func (this *PeekingIterator) next() int {
	if len(this.queue) == 0 {
		return this.iter.next()
	}
	next := this.queue[0]
	this.queue = this.queue[1:]
	return next
}

func (this *PeekingIterator) peek() int {
	if len(this.queue) != 0 {
		return this.queue[0]
	}
	peek := this.iter.next()
	this.queue = append(this.queue, peek)
	return peek
}
