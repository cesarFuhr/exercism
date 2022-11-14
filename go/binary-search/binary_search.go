package binarysearch

func SearchInts(list []int, key int) int {
	left, right := 0, len(list)-1
	for left <= right {
		mid := (left + right) / 2

		guess := list[mid]
		if guess == key {
			return mid
		}

		if guess > key {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

//func SearchInts(list []int, key int) int {
//	return recursive(list, key, 0, len(list)-1)
//}
//
//func recursive(list []int, key, left, right int) int {
//	mid := (left + right) / 2
//
//	if left > right {
//		return -1
//	}
//
//	guess := list[mid]
//
//	if guess < key {
//		return recursive(list, key, mid+1, right)
//	}
//
//	if guess > key {
//		return recursive(list, key, left, mid-1)
//	}
//
//	return mid
//}
