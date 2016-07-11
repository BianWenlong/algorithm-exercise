package sort

func ShellSort(data []int) {
	n:=len(data)
	k:=n/2
	for k>=1 {
		for i:=0;i<n;i++{
			for j:=i ;j>k-1;j-=k{
			if data[j]<data[j-k] {
				data[j],data[j-k]=data[j-k],data[j]
			}else{
				break
			}
		}
		}
		k/=2
	}
}
