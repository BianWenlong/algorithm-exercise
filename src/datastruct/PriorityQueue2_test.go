package datastruct

import (
	"testing"
	)

func TestPriorityQueue2(t *testing.T){
	queue:=NewPriorityQueueHeap(make([]Comparable,0))
	var a,b,c,d Integer
	a,b,c,d=1,2,3,4
	queue.Push(b)
	queue.Push(a)
	queue.Push(d)
	queue.Push(c)
	result:=queue.Pop()
	flag,_:=result.Compare(d)
	if flag!=0{
		t.Errorf("error")
	}
}