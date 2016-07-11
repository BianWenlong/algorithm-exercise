package datastruct

import "testing"

func PriorityQueuePushTest(t *testing.T){
	queue:=NewPriorityQueue()
	var a,b,c,d Integer
	a,b,c,d=1,2,3,4
	queue.Push(b)
	queue.Push(a)
	queue.Push(d)
	queue.Push(c)
	n:=queue.Pop()
	if n!=4{
		t.Errorf("error!")
	}
}

func NewPriorityQueue() *PriorityQueue{
	return &PriorityQueue{&Stack{&LinkList{nil,nil}},&Stack{&LinkList{nil,nil}}}
}

type Integer int 

func (a Integer) Compare(b Integer) int{
	if a<b{
		return -1
	}else if a==b {
		return 0
	}else {
		return 1
	}
}