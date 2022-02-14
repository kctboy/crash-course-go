package main

import "fmt"

func main() {
	//emails := make(map[string]string)

	//emails["Bob"] = "bob@gmail.com"
	//emails["kees"] = "kees@gmail.com"
	//emails ["klaas"] = "klaas@gmail.com"

	emails := map[string]string{"Bob": "bob@gmail.com", "kees": "kees@gmail.com", "klaas": "klaas@gmail.com"}
	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	delete(emails, "Bob")
	fmt.Println(emails)

}
