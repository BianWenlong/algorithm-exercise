package sort

func InsertSort(data []int ) {
	for i,_:=range data{
		for j:=i ;j>0;j--{
			if data[j]<data[j-1] {
				data[j],data[j-1]=data[j-1],data[j]
			}else{
				break
			}
		}
	}
}
