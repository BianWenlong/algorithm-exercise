package datastruct

type PriorityQueueHeap struct {
	data []Comparable
}

func (queue *PriorityQueueHeap)Push(node Comparable) bool{
	if len(queue.data)==0{
		queue.data=append(queue.data,nil)
	}
	queue.data=append(queue.data,node)
	location:=len(queue.data)-1
	for ;location>1;location/=2{
		flag,_:=queue.data[location].Compare(queue.data[location/2])
		if flag==1{
			queue.data[location],queue.data[location/2]=queue.data[location/2],queue.data[location]
		}else{
			break
		}
	}
	return true
}

func (queue *PriorityQueueHeap)Pop() Comparable{
	data:=queue.data
	l:=len(data)
	if l==1 || l==0{
		return nil
	}
	result:=data[1]
	data[1]=data[l-1]
	data=data[:l-1]
	for i:=1;i<=(l-2)/2;{
		var max Comparable
		maxLoc:=0
		flag,_:=data[i*2].Compare(data[i*2+1])
		if flag==1{
			max=data[i*2]
			maxLoc=i*2
		}else{
			max=data[i*2+1]
			maxLoc=i*2+1
		}
		flag,_=max.Compare(data[i])
		if flag==1{
			data[i],data[maxLoc]=data[maxLoc],data[i]
			i=maxLoc
		}else{
			break
		}
	}
	queue.data=data
	return result
}

func NewPriorityQueueHeap(data []Comparable) *PriorityQueueHeap{
	return &PriorityQueueHeap{data}
}