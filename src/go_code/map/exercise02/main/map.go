package main

import "fmt"

func main() {
	cities := make(map[string]string)
	cities["no1"] = "Kuala Lumpur"
	cities["no2"] = "Johor Bahru"
	cities["no3"] = "Pulau Pinang"

	fmt.Println(cities)

	//	Edit or Add City
	editOrAddCity(cities, "no1", "Putrajaya")
	editOrAddCity(cities, "no4", "Seremban")

	fmt.Println(cities)

	//	Delete
	delete(cities, "no4")

	fmt.Println(cities)

	//	Search
	val, ok := cities["no1"]
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("City not found")
	}

	fmt.Println("* * * * * * * * * * * * *")
	//	Iterate all cities
	for key, city := range cities {
		fmt.Println("Key:", key, "Value:", city)
	}
}

func editOrAddCity(cities map[string]string, city, cityNewName string) {
	cities[city] = cityNewName
}
