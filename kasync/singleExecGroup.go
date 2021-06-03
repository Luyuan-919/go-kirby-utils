package kasync

import "sync"

type call struct {
	sync.WaitGroup
	val interface{}
	err error
}

type SingleExecGroup struct {
	sync.Mutex
	callMap map[string]*call
}

func (s *SingleExecGroup) Do(key string,fn func()(interface{},error)) (interface{},error)  {
	s.Lock()
	if s.callMap == nil {
		s.callMap = make(map[string]*call)
	}
	if c, ok := s.callMap[key]; ok {
		s.Unlock()
		c.Wait()
		return c.val,c.err
	}
	c := new(call)
	c.Add(1)
	s.callMap[key] = c
	s.Unlock()

	func(){
		defer func() {
			if err := recover();err != nil {
			}
		}()
		c.val,c.err = fn()
	}()
	c.Done()
	s.Lock()
	delete(s.callMap,key)
	s.Unlock()

	return c.val,c.err
}
