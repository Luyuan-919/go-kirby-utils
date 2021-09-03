package ksort


func QuickSort(arr []int)   {
	if len(arr) == 0 {
		return
	}
	quickSort(arr,0,len(arr) - 1)
}

func quickSort(arr []int, start, end int) {
	if start > end {
		return
	}
	mid := arr[start]
	left,right := start,end
	for left < right {
		for right > left && arr[right] >= mid {
			right--
		}
		arr[left],arr[right] = arr[right],arr[left]
		for left < right && arr[left] <= mid {
			left++
		}
		arr[left],arr[right] = arr[right],arr[left]
	}
	quickSort(arr,start,left - 1)
	quickSort(arr,right+1,end)
}