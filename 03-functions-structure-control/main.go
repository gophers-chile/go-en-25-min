package main

import (
	"fmt"

	f "github.com/gophers-chile/go-en-25-min-ex3/funciones"
)

func main() {
	// Llamada a funciones
	fmt.Println("-----")
	f.ImprimeMensaje()
	fmt.Println("-----")
	f.ImprimeMensajeConArgumentos("¡Hola desde la función ImprimeMensajeConArgumentos!")
	fmt.Println("-----")
	mensaje := f.ImprimeMensajesConRetorno("¡Hola desde la función ImprimeMensajesConRetorno!")
	fmt.Println("Mensaje retornado:", mensaje)
	fmt.Println("-----")
	mensaje1, mensaje2 := f.ImprimeMensajesConMultipleRetornos("¡Hola!", "¡Adiós!")
	fmt.Println("Mensajes retornados:", mensaje1, "y", mensaje2)
	fmt.Println("-----")
	res, err := f.ImprimeMensajeCondicional("¡Hola desde la función ImprimeMensajeCondicional!", false)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Mensaje retornado:", res)
	fmt.Println("-----")
}
