package arrsli

import (
	"fmt"
)

func ImprimeArrays() {
	var arrayInt1 [3]int8
	arrayInt1[0] = 10
	arrayInt1[1] = 20
	arrayInt1[2] = 30
	fmt.Println("Array de enteros:", arrayInt1)

	var arrayString1 [4]string = [4]string{"Nico", "Paz", "Cris", "Magdalena"}
	fmt.Println("Array de strings:", arrayString1)
}
