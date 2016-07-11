package datastruct

import "testing"

func TestStackPush(t *testing.T){
	stack:=&Stack{&LinkList{nil,nil}}
	stack.Push(1)
	stack.Push(2)
	if stack.Size()!=2 {
		t.Errorf("error !")
	}
}

func TestStackPop(t *testing.T){
	stack:=&Stack{&LinkList{nil,nil}}
	stack.Push(1)
	stack.Push(2)
	data:=stack.Pop()
	if data!=2 {
		t.Errorf("error !")
	}
}