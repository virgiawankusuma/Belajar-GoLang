package main

import "fmt"

func main() {

	// membuat variable secara general
	var name string
	name = "virgiawan teguh"
	fmt.Println(name)
	fmt.Println("\n")

	name = "virgiawan kusuma"
	fmt.Println(name)
	fmt.Println("\n")

	// persingkat variable #1
	var umur = 21
	fmt.Println(umur)
	fmt.Println("\n")

	// persingkat variable #2
	alamat := "jepara, kuwasen"
	fmt.Println(alamat)
	fmt.Println("\n")

	// multidimensional variable
	var (
		firstName = "virgiawan"
		midName   = "teguh"
		lastName  = "kusuma"
	)

	fmt.Println(firstName, midName, lastName)
	fmt.Println("\n")
}
