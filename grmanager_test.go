package grtm

import (
	"fmt"
	"testing"
	"time"
)

func HelloWorld() {
	fmt.Println("Hello world!")
}

func Test_StopLoopGoroutine(t *testing.T) {
	gm := NewGrManager()
	// start a loop goroutine named 'test'
	gm.NewLoopGoroutine("test", func(args ...interface{}) {
		c := args[0].([]interface{})[0].(int) + args[0].([]interface{})[1].(int)
		fmt.Println(c)
	}, 1, 2)
	// start a loop goroutine named 'hello'
	gm.NewLoopGoroutine("hello", HelloWorld)
	time.Sleep(time.Duration(2) * time.Second)
	// stop a loop goroutine named 'test'
	err := gm.StopLoopGoroutine("test")
	if err != nil {
		t.Error(err)
	}
	time.Sleep(time.Duration(5) * time.Second)
}
