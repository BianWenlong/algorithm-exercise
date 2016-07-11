package datastruct

import (
		"testing"
		)

func TestLinklist(t *testing.T){
	list:=&LinkList{nil,nil}
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	if !list.Contains(1){
		t.Errorf("error !!!")
	}
}

func TestSize(t *testing.T){
	list:=&LinkList{nil,nil}
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	if list.Size()!=3{
		t.Errorf("error !!!")
	}
}

