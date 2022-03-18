package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	chanAny := make(chan interface{}, 3)
	chanAny <- 10
	chanAny <- "tom jacky"
	chanAny <- Cat{
		Name: "Tom",
		Age:  4,
	}
	<-chanAny
	<-chanAny
	cat := <-chanAny
	c, ok := cat.(Cat)
	fmt.Println(ok, c.Name)
}
