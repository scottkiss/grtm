package main

import (
	"fmt"
	"github.com/scottkiss/grtm"
	"time"
)

func normal() {
	fmt.Println("i am normal goroutine")
}

func funcWithParams(args ...interface{}) {
	fmt.Println(args[0].([]interface{})[0].(string))
	fmt.Println(args[0].([]interface{})[1].(string))
}

func main() {
	gm := grtm.NewGrManager()
	gm.NewGoroutine("normal", normal)
	fmt.Println("main function")
	gm.NewGoroutine("funcWithParams", funcWithParams, "hello", "world")
	time.Sleep(time.Second * time.Duration(5))
}
