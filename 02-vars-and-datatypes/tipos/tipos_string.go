package tipos

import (
	"fmt"
	"unicode/utf8"
)

func ImprimeTiposString() {
	fmt.Println("string :", "Hola a todos!")
	fmt.Println("string :", "Nico"+" "+" Avila "+" "+"Araya")
	fmt.Println("string :", "Nico\nAvila\nAraya")
	fmt.Println("string :", `Maria
Paz
Bustos`)
	fmt.Println("string length:", len("Nico"))
	fmt.Println("string length?:", len("Ŷ"))
	fmt.Println("string length!:", utf8.RuneCountInString("Ŷ"))
}
