package kpool
//
//import (
//	"errors"
//	"time"
//)
//
//
//const (
//	defQueueSize int32 = 2048
//	defMaxWorkers int32 = 1024
//	defMaxCoreWorkers = defMaxWorkers >> 1
//	defWorkerDeadTime = time.Second * 10
//	defCoreWorkerDeadTime = time.Minute * 10
//)
//
//type Executor interface {
//	Submit(task func()) error
//	Stop() error
//}
//
//var UnableToAddErr = errors.New("the queue is full and cannot be added")
//
//type WorkerPool struct {
//	maxCoreWorkers int32
//	maxWorkers int32
//	runningWorkers int32
//	workerQueue chan *func()
//}
//
//func NewDefPool() Executor {
//	return newPool(defMaxWorkers,defQueueSize)
//}
//
//func newPool(maxWorkers,queueSize int32) Executor {
//	if maxWorkers < 1 {
//		maxWorkers = defMaxWorkers >> 2
//	}
//	if maxWorkers > defMaxWorkers {
//		maxWorkers = defMaxWorkers
//	}
//	if queueSize < 1 {
//		queueSize = defQueueSize >> 3
//	}
//	if queueSize > defQueueSize {
//		queueSize = defQueueSize
//	}
//	pool := &WorkerPool{
//		maxCoreWorkers: maxWorkers >> 1,
//		maxWorkers: maxWorkers,
//		runningWorkers: 0,
//		workerQueue: make(chan *func(),queueSize),
//	}
//	go pool.dispatch()
//	return pool
//}
//
//func (p *WorkerPool) Submit(task func())error  {
//	return nil
//}
//
//func (p *WorkerPool) Stop() error {
//	return nil
//}
//
//
//
//func (p *WorkerPool) dispatch() {
//
//}