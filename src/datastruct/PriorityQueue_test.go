package datastruct

import (
		"testing"
		)

func TestPriorityQueuePush(t *testing.T){
	queue:=NewPriorityQueue()
	var a,b,c,d Integer
	a,b,c,d=1,2,3,4
	queue.Push(b)
	queue.Push(a)
	queue.Push(d)
	queue.Push(c)
	n:=queue.Pop()
	n=queue.Pop()
	n=queue.Pop()
	n=queue.Pop()
	result,_:=n.Compare(a)
	if result!=0{
		t.Errorf("error!")
	}
}

func NewPriorityQueue() *PriorityQueue{
	return &PriorityQueue{&Stack{&LinkList{nil,nil}},&Stack{&LinkList{nil,nil}}}
}

type Integer int 

func (a Integer) Compare(c interface{}) (int,error){
	b:=c.(Integer)
	if a<b{
		return -1,nil
	}else if a==b {
		return 0,nil
	}else {
		return 1,nil
	}
}