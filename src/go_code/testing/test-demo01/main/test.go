package main

func addUpper(n int) (res int) {
	for i := 1; i <= n; i++ {
		res += i
	}
	return
}

func main() {
	//	Traditional way to test a function

	//res := addUpper(10)
	//if res != 55 {
	//	fmt.Println("False")
	//} else {
	//	fmt.Println("True")
	//}

	//	1.	Not convenient
	//	2.	Not Easy to manage
}
