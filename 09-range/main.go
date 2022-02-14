package main

import "fmt"

func main() {
	ids := []int{33, 43, 64, 11, 23, 16}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// range met map

	emails := map[string]string{"Bob": "bob@gmail.com", "kees": "kees@gmail.com", "klaas": "klaas@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Naam: " + k)
	}
}
