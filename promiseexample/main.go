package main

import (
	"fmt"
	"syscall/js"
	"time"

	"github.com/realPy/hogosuru"
	"github.com/realPy/hogosuru/promise"
)

func main() {
	hogosuru.Init()

	fmt.Printf("p1 is launch\n")
	p1, _ := promise.New(func(resolvefunc, errfunc js.Value) (interface{}, error) {
		fmt.Printf("p1 is calculate\n")
		time.Sleep(time.Duration(2) * time.Second)
		fmt.Printf("p1 end calculate\n")
		return js.ValueOf("p1 return"), nil
	})

	fmt.Printf("p1 is in progress....\n")

	p1.Then(func(i interface{}) *promise.Promise {
		if str, ok := i.(string); ok {

			fmt.Printf("p1 is finished with %s\n", str)
		}

		return nil
	}, func(e error) {

		fmt.Printf("There was an error %s\n", e.Error())
	})

	ch := make(chan struct{})
	<-ch
}
