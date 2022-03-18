package main

import (
	"fmt"
	"sync"
)

var map1 = make(map[int]int, 200)

//	Global mutexes =>>	全局的互斥锁
var lock sync.Mutex

//	There are conflict between the map variable
func main() {
	for i := 1; i <= 20; i++ {
		go times(map1, i)
	}

	//time.Sleep(time.Second)
	lock.Lock()
	for i, v := range map1 {
		fmt.Printf("map[%v]=%d\n", i, v)
	}
	lock.Unlock()
}

func times(map1 map[int]int, n int) {
	sum := 1
	for i := 1; i <= n; i++ {
		sum *= i
	}
	lock.Lock()
	map1[n] = sum
	lock.Unlock()
}
