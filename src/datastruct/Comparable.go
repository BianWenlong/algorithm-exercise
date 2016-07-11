package datastruct

type Comparable interface{
	Compare(other interface{}) (n int,err error)
}