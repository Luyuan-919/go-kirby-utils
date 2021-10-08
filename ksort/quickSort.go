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
	//这里使用<而不是<= 的原因：如果有等于，会一直在left == right的死循环里进行交换
	//for循环里的思路：
	//先找一个数应该在的位置，经过一次循环后，right后的数字都比mid大，left前的数字都比mid小，而left到right里还是不确定的，因此在这个区间里继续进行排序
	//直到left == right，此时mid一定在该在的位置上。
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