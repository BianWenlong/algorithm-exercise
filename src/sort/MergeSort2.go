package sort 

func MergeSort(data []int ){
	l:=len(data)
	for i:=1 ;i<l ;i*=2{
		for j:=0;j<=l-i;j+=2*i{
			Merge(j,j+i,min(l,j+2*i),data)
		}
	}
}

func Merge(beg,mid,end int ,data []int) {
	if(beg>=end){
		return 
	}
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

func min(a,b int) int {
	if a<b {
		return a
	}
	return b
}
