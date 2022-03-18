package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println("CPU number =", cpuNum)

	//	Can set Max Processor that are going to use
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("OK")
}
