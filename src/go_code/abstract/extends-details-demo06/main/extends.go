package main

import "fmt"

type A struct {
	Name string
	Age  int
}

type B struct {
	Name string
	Age  int
}

type C struct {
	A
	B
}

type D struct {
	a A //	Named struct
}

type Goods struct {
	Name  string
	Price float64
}

type Brand struct {
	Name    string
	Address string
}

type Television struct {
	*Goods
	*Brand
	int
}

func main() {
	var c C

	//ambiguous selector c.Name
	//c.Name = "Chloe Gan"
	//fmt.Println(c.Name)

	c.A.Name = "Chloe Gan"
	c.B.Name = "Ezra Chai"
	fmt.Println("c.A.Name =", c.A.Name)
	fmt.Println("c.B.Name =", c.B.Name)

	fmt.Println("*******************")

	var d D
	d.a.Name = "Chloe Gan"
	fmt.Println(d.a.Name)

	tv := &Television{
		Goods: &Goods{
			Name:  "Z9J | BRAVIA XR | MASTER Series| Full Array LED | 8K | High Dynamic Range (HDR) | Smart TV (Google TV)",
			Price: 73999,
		},
		Brand: &Brand{
			Name:    "Sony",
			Address: "Damansara",
		},
	}
	fmt.Println(*tv.Goods, *tv.Brand)
	tv.int = 123
	fmt.Println(tv.int)
}
