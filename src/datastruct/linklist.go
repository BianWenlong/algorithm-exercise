package datastruct

type Node struct {
	Data interface{}
	Next *Node
}

type LinkList struct{
	Last *Node
	First *Node
}

type NilError struct{
	who string
}

func (e *NilError) Error() string{
	return e.who
}

func (list *LinkList)Insert(data interface{}) (interface{},error){
	if data==nil{
		return data,&NilError{"data == nil"}
	}
	node:=Node{data,nil}
	if list.First==nil {
		list.First,list.Last=&node,&node 
	}else{
		node.Next,list.First=list.First,&node
	}
	return data,nil
}

func (list *LinkList) GetAndDelFirst() interface{}{
	if(list.First==nil){
		return nil
	}
	data:=list.First.Data
	list.First=list.First.Next
	return data
}
func (list *LinkList)Contains(data interface{}) bool{
	if(data==nil){
		return false
	}
	n:=list.First
	for{
		if n==nil {
			break
		}else if data==n.Data {
			return true
		}else{
			n=n.Next
		}
	}
	return false
}

func (list *LinkList) Size() int{
	i,n:=0,list.First
	for{
		if(n==nil){
			break
		}else{
			n=n.Next
			i++
		}
	}
	return i
}