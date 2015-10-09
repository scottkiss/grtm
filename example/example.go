package main

import (
	"fmt"
	"github.com/scottkiss/grtm"
	"time"
)

func normal() {
	fmt.Println("i am normal goroutine")
}

func main() {
	gm := grtm.NewGrManager()
	gm.NewGoroutine("normal", normal)
	fmt.Println("main function")
	time.Sleep(time.Second * time.Duration(5))
}
