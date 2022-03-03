//hello.go is in main package, in go every file must in a package
package main

//import a package called "fmt",you can use "fmt" package's function after import the package, ex: fmt.Println
import "fmt"

// func is a keyword, "main()" is a main function
func main(){

	//Calls the Println function in the "fmt" package
	fmt.Println("Hello World")
}