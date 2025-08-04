package common

// returns idx of largest number lesser than toFind
func BinarySearchIdxLesserThan(toFind int, ary []int) int {
	start := 0
	end := len(ary) - 1
	for start <= end {
		mid := (start + end) / 2
		if toFind == ary[mid] {
			return mid
		}

		if toFind > ary[mid] {
			if mid == len(ary)-1 || toFind < ary[mid+1] {
				return mid
			}

			start = mid + 1
		} else {
			if mid == 0 || toFind > ary[mid-1] {
				return mid - 1
			}

			end = mid - 1
		}
	}

	return -1
}

// for comparator func(T1, T2):
//
//	0 means equal, 1 means t1 > t2, -1 means t1 < t2
func BinarySearchIdxLesserThanDuplicatesTypeGeneric[T comparable](toFind T, ary []T, comparator func(T, T) int) []int {
	start := 0
	end := len(ary) - 1
	for start <= end {
		mid := (start + end) / 2
		if comparator(toFind, ary[mid]) == 0 {
			matchStart := mid
			matchEnd := mid
			for matchStart > 0 && comparator(ary[matchStart], ary[matchStart-1]) == 0 {
				matchStart--
			}

			for matchEnd < len(ary)-1 && comparator(ary[matchEnd], ary[matchEnd+1]) == 0 {
				matchEnd++
			}

			return []int{matchStart, matchEnd}
		}

		if comparator(toFind, ary[mid]) == 1 {
			if mid == len(ary)-1 || comparator(toFind, ary[mid+1]) == -1 {
				return []int{mid, mid}
			}

			start = mid + 1
		} else {
			if mid == 0 || comparator(toFind, ary[mid-1]) == 1 {
				return []int{mid - 1, mid - 1}
			}

			end = mid - 1
		}
	}

	return []int{}
}
