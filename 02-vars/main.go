package main

import "fmt"

func main() {
	var name string = "brad"
	var age = 37
	const isCool = true
	size := 1.3

	email, home := "stb@dds.nl", "oploo"
	fmt.Println(name, age, isCool, size, email, home)
	fmt.Printf("%T\n", size)

}
