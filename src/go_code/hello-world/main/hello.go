//hello.go is in main package, in Go every file must in a package
package main

//import a package called "fmt",you can use "fmt" package's function after import the package, ex: fmt.Println
import "fmt"

// func is a keyword, "main()" is a main function
func main(){

	//Calls the Println function in the "fmt" package
	fmt.Println("Hello world!")

	//go build -o Hallo hello.go -> Hallo.exe
	//go build hello.go -> hello.exe
}