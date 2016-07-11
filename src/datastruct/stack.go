package datastruct

type Stack struct{
	list *LinkList
}

func (stack *Stack) Push (data interface{}) bool{
	_,e:=stack.list.Insert(data)
	if(e!=nil){
		return false
	}
	return true
}

func (stack *Stack) Pop() interface{} {
	return stack.list.GetAndDelFirst()
}

func (stack *Stack) Size() int {
	return stack.list.Size()
}
