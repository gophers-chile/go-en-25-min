package main

import (
	"fmt"

	t "github.com/gophers-chile/go-en-25-min-ex2/tipos"
)

func main() {

	// Declaracion de variables
	var nombre string = "Nico"
	fmt.Println("Nombre:", nombre)
	apellido := "Avila"
	fmt.Println("Apellido:", apellido)
	var hijo, hija string = "Cristobal", "Magdalena"
	fmt.Println("Hijo:", hijo, "Hija:", hija)

	// Llama a funcion en paquete de tipos numericos
	fmt.Println("-----")
	t.ImprimeTiposNumericos()
	fmt.Println("-----")
	t.ImprimeTiposString()
}
