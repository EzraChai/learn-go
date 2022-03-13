package main

import "fmt"

func main() {

	users := make(map[string]map[string]string)

	users["chloegan"] = make(map[string]string, 2)
	users["chloegan"]["nickname"] = "Chloe Gan"
	users["chloegan"]["pwd"] = "chloegan0920"

	modifyUser(users, "ezrachai")
	fmt.Println(users)

}

func modifyUser(users map[string]map[string]string, name string) {
	_, b := users[name]
	if b {
		users[name]["pwd"] = "888888"
	} else {
		users[name] = make(map[string]string, 2)
		users[name]["nickname"] = ""
		users[name]["pwd"] = ""
	}
}
