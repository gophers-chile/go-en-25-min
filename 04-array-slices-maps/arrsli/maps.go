package arrsli

import (
	"fmt"
)

func ImprimeMaps() {
	// Definición e inicialización de un map
	var mapComunidades = map[string]string{
		"Comunidad1": "Golang CL",
		"Comunidad2": "Python Chile",
		"Comunidad3": "UX Rangel",
	}
	fmt.Println("Map de comunidades:", mapComunidades)

	// Comprueba si el elemento existe en el map
	var comunidad, ok = mapComunidades["Comunidad4"]
	if ok {
		fmt.Println("Comunidad encontrada:", comunidad)
	} else {
		fmt.Println("Comunidad no encontrada")
	}
}
