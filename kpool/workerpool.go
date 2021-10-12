package kpool

import (
	"errors"
	"fmt"
	"sync/atomic"
	"time"
)

type Status int32
type TaskStatus int32

const (
	StatusRunning  = 1
	StatusDispatch = 2
	StatusClose = -1

	TaskStatusInit = -1
	TaskStatusFinish  = 1

	defWorkerIdleTime = time.Second * 8
	defQueueSize = 4096

	MaxIdleTime = time.Second * 15
	MaxWorkers  =  8192
	MaxQueueSize = 8192 << 8
)

var (
	UseClosedPoolErr = errors.New("you are using a closed pool")
	UnableToAddErr = errors.New("the queue is full and cannot be added")
)

type workerPool struct {
	//最大协程数量
	maxWorkers int32
	//正在运行的工人数量
	runningWorkers int32
	//协程等待的最大时间
	maxWorkerIdleTime time.Duration
	//协程池的状态
	status Status
	//任务队列
	taskQueue chan *func()
}

//任务 当有一些需要返回值的任务时使用这个
type Task struct {
	fn func() (interface{},error)
	retVal interface{}
	retErr error
	status TaskStatus
}

func NewTask(fn func() (interface{},error))  *Task{
	return &Task{
		fn:     fn,
		retVal: nil,
		retErr: nil,
		status: TaskStatusInit,
	}
}

func NewWorkerPool() *workerPool {
	return newPool(MaxWorkers,defQueueSize,defWorkerIdleTime)
}

func NewWorkPoolWithMaxWorkers(workers int32) *workerPool {
	return newPool(workers,defQueueSize,defWorkerIdleTime)
}

func NewWorkPoolWithMaxWorkersAndQueueSize(workers,queueSize int32) *workerPool {
	return newPool(workers,queueSize,defWorkerIdleTime)
}

func NewWorkPoolWithMaxWorkersAndQueueSizeAndIdleTime(workers,queueSize int32,idleTime time.Duration) *workerPool {
	return newPool(workers,queueSize,idleTime)
}


//初始化协程池
func newPool(maxWorkers, QueueSize int32, IdleTime time.Duration) *workerPool {
	//对传入的参数进行修正
	if maxWorkers < 1 || maxWorkers > MaxWorkers {
		maxWorkers = MaxWorkers
	}
	if QueueSize < 1 || QueueSize > MaxQueueSize {
		QueueSize = defQueueSize
	}
	if IdleTime > MaxIdleTime {
		IdleTime = defWorkerIdleTime
	}
	pool := &workerPool{
		maxWorkers: maxWorkers,
		runningWorkers: 0,
		maxWorkerIdleTime: IdleTime,
		status:StatusRunning,
		taskQueue: make(chan *func(),QueueSize),
	}
	go pool.dispatch()
	return pool
}

func (p *workerPool) dispatch()  {
Loop:
	for atomic.LoadInt32((*int32)(&p.status)) == StatusRunning &&
	atomic.LoadInt32(&p.runningWorkers) < p.maxWorkers{
		select {
		case tk, ok := <-p.taskQueue:
			if !ok {
				break Loop
			}
			atomic.AddInt32(&p.runningWorkers, 1)
			go p.worker(tk)
		}
	}
}

func (p *workerPool) worker(tk *func()) {
	(*tk)()
	if p.doWorker() {
		atomic.AddInt32(&p.runningWorkers, -1)
	}
}

func (p *workerPool) doWorker() bool {
	if p.maxWorkerIdleTime > 0 {
		//开启计时器，该计时器用于控制工作协程的过期时间，每次运行完task任务后计时器会被重置
		idle := time.NewTimer(p.maxWorkerIdleTime)
		//协程池没有关闭的时候都可以运行
		for atomic.LoadInt32((*int32)(&p.status)) != int32(StatusClose) {
			select {
			case task, ok := <-p.taskQueue:
				if !ok {
					break
				}
				(*task)()
				idle.Reset(p.maxWorkerIdleTime)
			//当正在运行的协程长时间没有获取到任务超时后，会将该协程改变为任务分发协程
			//同时会将工作协程减一，同理当该协程在获取到任务，再次将工作协程数开启到最大数量时会退出
			case <-idle.C:
				if atomic.LoadInt32(&p.runningWorkers) <= p.maxWorkers-1 &&
					atomic.CompareAndSwapInt32((*int32)(&p.status),
						int32(StatusRunning), int32(StatusDispatch)) {
					atomic.AddInt32(&p.runningWorkers, -1)
					p.dispatch()
					return false
				}
				break
			}
		}
	} else {
		//如果没有设置超时时间，当任务完成后，则自动工作协程自动退出
		for atomic.LoadInt32((*int32)(&p.status)) != int32(StatusClose) {
			task, ok := <-p.taskQueue
			if !ok {
				break
			}
			(*task)()
		}
	}
	return true
}


// Submit 提交任务
func (p *workerPool) Submit(task func()) error {
	//如果任务为空则直接返回
	if task == nil {
		return nil
	}
	//当协程池的状态为关闭相关时，抛出异常
	if  p.status == StatusClose  {
		return UseClosedPoolErr
	}
	//将任务放入任务队列中，准备执行
	select {
	case p.taskQueue <- &task:
	default:
		return UnableToAddErr
	}
	return nil
}

// SubmitWait 等待任务完成的提交
func (p *workerPool) SubmitWait(task func()) error {
	//如果任务为空则直接返回
	if task == nil {
		return nil
	}
	//如果协程池状态为关闭相关状态则抛出异常
	if p.status == StatusClose {
		return UseClosedPoolErr
	}
	//构造一个提交完成管道，同时将需要执行的任务封装为doneFunc方法
	doneChan := make(chan bool)
	var doneFunc = func() {
		task()
		close(doneChan)
	}
	//阻塞等待，当被封装成doneFunc的任务执行完成并关闭管道时，会返回
	select {
	case p.taskQueue <- &doneFunc:
		<-doneChan
	default:
		return UnableToAddErr
	}
	return nil
}

// SubmitTask 提交任务
func (p *workerPool) SubmitTask(task *Task) error {
	//如果任务为空则直接返回
	if task == nil {
		return nil
	}
	//当协程池的状态为关闭相关时，抛出异常
	if  p.status == StatusClose {
		return UseClosedPoolErr
	}
	fn := func() {
		task.status = TaskStatusInit
		task.retVal, task.retErr = task.fn()
		task.status = TaskStatusFinish
	}
	//将任务放入任务队列中，准备执行
	select {
	case p.taskQueue <- &fn:
	default:
		return UnableToAddErr
	}
	return nil
}

func (p *workerPool) Stop()  {
	p.stop()
	return
}

func (p *workerPool) stop()  {
	if atomic.LoadInt32((*int32)(&p.status)) == int32(StatusClose) {
		return
	}
	atomic.StoreInt32((*int32)(&p.status), int32(StatusClose))
	for len(p.taskQueue) > 0{
		time.Sleep(1e6)
		fmt.Println("-1")
	}
	close(p.taskQueue)
}