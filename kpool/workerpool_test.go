package kpool

import (
	"sync/atomic"
	"testing"
)

func BenchmarkWorkerPool_Submit2(b *testing.B) {
	var num1,num2 int32
	num1 = 345645
	num2 = 465343
	flag := false
	for i := 0; i < b.N; i++{
		x := func() {
			if flag {
				atomic.StoreInt32(&num1,num2)
				flag = false
			}else{
				atomic.StoreInt32(&num2,num1)
				flag = true
			}
		}
		x()
	}
}

func BenchmarkWorkerPool_Submit(b *testing.B) {
	var num1,num2 int32
	num1 = 345645
	num2 = 465343
	flag := false
	pool := NewWorkerPool()
	for i := 0; i < b.N; i++{
		x := func() {
			if flag {
				atomic.StoreInt32(&num1,num2)
				flag = false
			}else{
				atomic.StoreInt32(&num2,num1)
				flag = true
			}
		}
		_ = pool.Submit(x)
	}
}
