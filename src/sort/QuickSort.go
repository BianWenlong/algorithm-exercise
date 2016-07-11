package sort

func QuickSort(data []int) {
	l:=len(data)
	if(l==0 || l==1){
		return 
	}
	j:=-1
	flag:=data[l-1]
	for i,v:=range data{
		if(v>flag){
			continue
		}else{
			j++
			data[j],data[i]=v,data[j]
		}
	}
	QuickSort(data[:j])
	QuickSort(data[j:])

}
