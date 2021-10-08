package str

import "unsafe"

func BytesToString(bys []byte) string {
	if bys == nil || len(bys) == 0 {
		return ""
	}
	//获取bys的地址，先转为unsafe.Pointer，再把这个指针转为string类型的，再取值
	return *(*string)(unsafe.Pointer(&bys))
}

func StringToBytes(s *string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(s)) // 获取s的起始地址开始后的两个 uintptr 指针
	h := [3]uintptr{x[0], x[1], x[1]}     // 构造三个指针数组
	return *(*[]byte)(unsafe.Pointer(&h))
}