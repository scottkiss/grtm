package main

import (
	"fmt"
	"github.com/scottkiss/grtm"
	"time"
)

func myfunc() {
	fmt.Println("do something repeat by interval 4 seconds")
	time.Sleep(time.Second * time.Duration(4))
}

func main() {
	gm := grtm.NewGrManager()
	gm.NewLoopGoroutine("myfunc", myfunc)
	fmt.Println("main function")
	time.Sleep(time.Second * time.Duration(40))
	fmt.Println("stop myfunc goroutine")
	gm.StopLoopGoroutine("myfunc")
	time.Sleep(time.Second * time.Duration(80))
}
