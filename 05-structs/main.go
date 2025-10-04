package main

import (
	"fmt"

	st "github.com/gophers-chile/go-en-25-min-ex5/structs"
)

func main() {
	fmt.Println("Structs en Go")
	paciente := st.Diagnosticar1()
	fmt.Println(paciente)
}
