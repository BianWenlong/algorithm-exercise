package datastruct

type PriorityQueue struct{
	dataStack *Stack
	maxStack *Stack
}

func (queue *PriorityQueue) Push(data Comparable) bool {
	for n:=queue.dataStack.Pop();n!=nil;{
		if n.Compare(data) > 0 {
			queue.maxStack.Push(n)
			n=queue.dataStack.Pop()
		}else{
			queue.dataStack.Push(n)
			queue.maxStack.Push(data)
			break
		}
	}
	if queue.dataStack.Size()==0 {
		queue.dataStack.Push(data)
	}
	for n:=queue.maxStack.Pop(); n!=nil ;{
		queue.dataStack.Push(n)
		n=queu.dataStack.Pop()
	}

	return true
}

func (queue *PriorityQueue) Pop() Comparable{
	return queue.dataStack.Pop()
}

func (queue *PriorityQueue) Size() int {
	return queue.dataStack.Size()
}