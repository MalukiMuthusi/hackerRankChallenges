package insertionsort

func InsertionSort(a []int) []int{
	for j:=1; j < len(a); j++{
		key := a[j]

		i := j-1

		for ;i>=0 && a[i]>key; {
			a[i+1]=a[i]
			i=i-1
		}

		a[i+1]=key
	}
	return a
}