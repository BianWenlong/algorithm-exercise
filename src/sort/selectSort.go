package sort

func SelectSort( data []int ) {
	if len(data)==0{
		return
	}
	for i,d:=range data {
		flag,min:=i,d
		for j:=i+1;j<len(data);j++{
			if data[j]<min{
				flag,min=j,data[j]
			}
		}
		data[i],data[flag]=min,data[i]
	}
	return 
}
