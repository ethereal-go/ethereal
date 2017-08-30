package utils

type FuncQueue struct {
	queue map[interface{}][]interface{}
	/**
	/ example queue
	/ 				func ... =>  ... arguments
	*/
}

func (q FuncQueue) Start() FuncQueue {
	q.queue = map[interface{}][]interface{}{}
	return q
}

func (q FuncQueue) Then(operation func(...interface{}), arguments ...interface{}) FuncQueue {
	q.queue[operation] = arguments
	return q
}

func (q FuncQueue) Exec() {
	//for operation, arguments := range q.queue  {
	//	//operation.(func)(arguments)
	//}
}
