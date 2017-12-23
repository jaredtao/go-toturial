package main

import "fmt"

func main() {
	var personSalary map[string]int
	if personSalary == nil {
		personSalary = make(map[string]int)
	}

	imap := make(map[string]string)
	imap["ME"] = "Hello"
	imap["YOU"] = "World"
	imap["SHE"] = "SHASHA"
	fmt.Println(imap, len(imap))

	for key, value := range imap {
		fmt.Println(key, value)
	}

	delete(imap, "SHE")
	fmt.Println(imap, len(imap))
}
