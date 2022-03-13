package main

import "fmt"

func main() {
	studentInformation := make(map[string]map[string]string)
	studentInformation["01"] = make(map[string]string, 2)
	studentInformation["01"]["name"] = "Tom"
	studentInformation["01"]["gender"] = "man"

	studentInformation["02"] = make(map[string]string, 2)
	studentInformation["02"]["name"] = "Chloe"
	studentInformation["02"]["gender"] = "woman"

	studentInformation["03"] = make(map[string]string, 2)
	studentInformation["03"]["name"] = "Ezra"
	studentInformation["03"]["gender"] = "man"

	// O(n^2)
	for key, studentInfo := range studentInformation {
		fmt.Printf("[%v]", key)
		for key2, value := range studentInfo {
			fmt.Printf("%v:[%v] \t", key2, value)
		}
		fmt.Println()
	}

	fmt.Println("length:", len(studentInformation))
}
