package main

import "fmt"

func main() {

	//	SLICE OF MAP

	monsters := make([]map[string]string, 2)

	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "Chibai"
		monsters[0]["age"] = "500"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "Shiba"
		monsters[1]["age"] = "1000"
	}
	//if monsters[2] == nil {
	//	monsters[2] = make(map[string]string, 2)
	//	monsters[2]["name"] = "Fuck"
	//	monsters[2]["age"] = "10000"
	//}

	//	Append Fix This Problem
	newMonster := map[string]string{
		"name": "Fuck",
		"age":  "10000",
	}
	monsters = append(monsters, newMonster)

	fmt.Println(monsters)

}
