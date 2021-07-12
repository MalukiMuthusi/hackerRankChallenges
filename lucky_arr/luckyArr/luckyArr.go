package luckyArr

// LuckyArr returns true or false
//	a lucky array is an array that can be split into three parts with values of equal sums
func LuckyArr(a []int) bool {

	for i := range a {
		var lSum = 0
		for ii := 0; ii <= i; ii++ {
			lSum += a[ii]
		}

		var rSum = 0
		var ir = len(a) - 1
		for ; ir > i; ir-- {
			rSum += a[ir]
			if rSum == lSum {
				break
			}
		}

		var mSum = 0
		for im := i + 1; im < ir; im++ {
			mSum += a[im]
		}

		if lSum == mSum && lSum == rSum {
			return true
		}

	}
	return false
}
