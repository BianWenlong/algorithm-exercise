package sort

func MergeSort(data []int) {
	l:=len(data)
	if( l==0 || l==1){
		return
	}
	MergeSort(data[:l/2])
	MergeSort(data[l/2:])
	Merge(0,l/2,l,data)
}

func Merge(beg,mid,end int ,data []int) {
	i,j:=beg,mid
	result:=make([]int,0)
	for {
		if i>=mid || j>=end{
			break;
		}
		if(data[i]<data[j]){
			result=append(result,data[i])
			i++
		}else{
			result=append(result,data[j])
			j++
		}
	}
	if(i>=mid){
		result=append(result,data[j:end]...)
	}else{
		result=append(result,data[i:mid]...)
	}
	for i,v:=range result{
		data[i+beg]=v
	}
}
