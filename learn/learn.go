package main

import (
	"fmt"

	"sync"
	"sync/atomic"
	"time"
)

func main() {
	//testSlice()
	//println(testDefer())
	//casRun()
	runUnitPtr()
}

func runUnitPtr() {
	println(uintptr(1))
	println(4 << (^uintptr(0) >> 63))
}

func testSlice() {
	var a []int
	//a := make([]int, 0)
	//a := []int{}
	println(a)
}

func testDefer() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

var (
	counter int32          //计数器
	wg      sync.WaitGroup //信号量
)

func casRun() {
	threadNum := 5
	wg.Add(threadNum)
	for i := 0; i < threadNum; i++ {
		go incCounter(i)
	}
	wg.Wait()

}
func incCounter(index int) {
	defer wg.Done()

	spinNum := 0
	for {
		// 原子操作
		old := counter
		ok := atomic.CompareAndSwapInt32(&counter, old, old+1)
		if ok {
			break
		} else {
			spinNum++
		}
	}
	fmt.Printf("thread,%d,spinnum,%d\n", index, spinNum)
	time.Now()

	for _, ints := range [][2]int{{1, 2}} {
		println(ints)

	}
}
