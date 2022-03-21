package main

import (
	"syscall/js"

	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/promise"
)

func main() {
	hogosuru.Init()

	p1, _ := promise.New(func(resolvefunc, errfunc js.Value) (interface{}, error) {
		println("Begin p1")
		w1, _ := promise.SetTimeout(8000)
		w1.Then(func(i interface{}) *promise.Promise {
			println("End p1")
			resolvefunc.Invoke(js.ValueOf("resolve with p1"))
			return nil
		}, nil)

		return nil, nil

	})

	p2, _ := promise.New(func(resolvefunc, errfunc js.Value) (interface{}, error) {
		println("Begin p2")
		w1, _ := promise.SetTimeout(5000)
		w1.Then(func(i interface{}) *promise.Promise {
			println("End p2")
			resolvefunc.Invoke(js.ValueOf("resolve with p2"))
			return nil
		}, nil)

		return nil, nil

	})

	p3, _ := promise.All(p1, p2)

	p3.Then(func(i interface{}) *promise.Promise {
		println("All is finished!")

		return nil
	}, func(e error) {

	})

	ch := make(chan struct{})
	<-ch

}
