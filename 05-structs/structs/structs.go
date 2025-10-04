package structs

import "strconv"

type Paciente struct {
	Nombre   string
	Edad     int
	Sintomas [2]string
}

func Diagnosticar1() string {
	var pac Paciente = Paciente{
		Nombre:   "Juan Perez",
		Edad:     30,
		Sintomas: [2]string{"fiebre", "tos"},
	}
	return "El paciente " + pac.Nombre + " de edad " + strconv.Itoa(pac.Edad) + " presenta síntomas de: " + pac.Sintomas[0] + " y " + pac.Sintomas[1]
}
