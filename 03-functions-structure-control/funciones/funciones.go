package funciones

import (
	"fmt"

	"errors"
)

func ImprimeMensaje() {
	fmt.Println("¡Hola desde la función ImprimeMensaje!")
}

func ImprimeMensajeConArgumentos(mensaje string) {
	fmt.Println("Mensaje:", mensaje)
}

func ImprimeMensajesConRetorno(mensaje1 string) string {
	return mensaje1
}

func ImprimeMensajesConMultipleRetornos(mensaje1, mensaje2 string) (string, string) {
	return mensaje1, mensaje2
}

func ImprimeMensajeCondicional(mensaje string, imprimir bool) (string, error) {
	var err error
	if imprimir {
		return "Mensaje: " + mensaje, err
	} else {
		err = errors.New("la variable 'imprimir' es false, no se imprimió el mensaje")
		return "", err
	}
}
