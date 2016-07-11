package datastruct

type PriorityQueue struct{
	dataStack *Stack
	maxStack *Stack
}

func (queue *PriorityQueue) Push(data Comparable) bool {
	for n:=queue.dataStack.Pop();n!=nil;{
		c:=n.(Comparable)
		result,_:=c.Compare(data)
		if result > 0 {
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
		n=queue.maxStack.Pop()
	}

	return true
}

func (queue *PriorityQueue) Pop() Comparable{
	result:=queue.dataStack.Pop()
	if(result==nil){
		return nil
	}
	result1:=result.(Comparable)
	return result1
}

func (queue *PriorityQueue) Size() int {
	return queue.dataStack.Size()
}