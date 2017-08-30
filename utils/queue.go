package utils

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	ErrParamsNotAdapted = errors.New("The number of params is not adapted.")
)

type FuncQueue struct {
	q     map[string]reflect.Value
	queue map[interface{}][]interface{}
	source interface{}
	/**
	/ example queue
	/ 				func ... =>  ... arguments
	*/
}

func (q FuncQueue) start() FuncQueue {
	q.queue = map[interface{}][]interface{}{}
	return q
}

//func (q FuncQueue) Then(operation interface{}, arguments ...interface{}) FuncQueue {
//	q.queue[operation] = arguments
//	return q
//}

func (q FuncQueue) Source(val interface{}) FuncQueue {
	q.source = val
	return  q
}

/**
/ @link https://mikespook.com/2012/07/function-call-by-name-in-golang/ note
 / @link https://bitbucket.org/mikespook/golib/src/27c65cdf8a772c737c9f4d14c0099bb82ee7fa35/funcmap/funcmap.go?at=default&fileviewer=file-view-default
*/
func (q FuncQueue) Exec() (res interface{}, err error) {

	for operation, arguments := range q.queue {

		//operation.(func(...interface{})interface{})(arguments...)
		//operation.(func)(arguments)
		//f := reflect.ValueOf(operation)

		//if len(arguments) != operation.(reflect.Value).Type().NumIn() {
		//	err := errors.New("The number of params is not adapted.")
		//	return nil, err
		//}
		in := make([]reflect.Value, len(arguments))
		for k, param := range arguments {
			in[k] = reflect.ValueOf(param)
		}
		result := operation.(reflect.Value).Call(in)
		fmt.Println(result, "@@@")
		return result, err
	}
	return
}
func (q FuncQueue) Then(operation interface{}, arguments ...interface{}) FuncQueue {
	if q.queue == nil {
		q = q.start()
	}
	v := reflect.ValueOf(operation)
	v.Type().NumIn()
	q.queue[v] = arguments
	return q
}

func (f FuncQueue) bind(name string, fn interface{}) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(name + " is not callable.")
		}
	}()
	v := reflect.ValueOf(fn)
	v.Type().NumIn()
	f.q[name] = v
	return
}
func (f FuncQueue) Call(name string, params ...interface{}) (result []reflect.Value, err error) {
	if _, ok := f.q[name]; !ok {
		err = errors.New(name + " does not exist.")
		return
	}
	if len(params) != f.q[name].Type().NumIn() {
		err = ErrParamsNotAdapted
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.q[name].Call(in)
	return
}
