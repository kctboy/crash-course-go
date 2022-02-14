package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	// gebruik * om de waarde van adress te zien
	fmt.Println(*b)

	// change value pointer
	*b = 10

	fmt.Println(a)
}
