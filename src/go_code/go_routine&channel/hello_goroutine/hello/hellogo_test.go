package hello

import "testing"

func test() {
	for i := 0; i < 100000; i++ {
		HelloGo(i)
	}
}

//	MPG
//	M =>> Machine
//	P =>> Processor
//	G =>> Goroutine

func TestGoRoutine(t *testing.T) {
	//	without go
	//--- PASS: TestGoRoutine (20.01s)
	//test()

	//	with go
	//--- PASS: TestGoRoutine (10.00s)
	go test() //	Use Goroutine
	for i := 0; i < 100000; i++ {
		HelloWorld(i)
	}

}
