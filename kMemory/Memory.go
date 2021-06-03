package kMemory

import "unsafe"

func SizeofMemoryInt(value interface{}) (size int){
	s := unsafe.Sizeof(value)
	size = int(s)
	return
}
