package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Estructura para la respuesta JSON
type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// Handler para el healthcheck
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := Response{
		Message: "La API está funcionando correctamente",
		Status:  "ok",
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	// Cre un servidor HTTP
	mux := http.NewServeMux()

	// Registra los handlers
	mux.HandleFunc("/health", healthHandler)

	// Configuracion del servidor
	port := ":8080"
	fmt.Printf("Inicializando servidor en el puerto %s\n", port)
	fmt.Println("Endpoints disponibles:")
	fmt.Println("- GET http://localhost:8080/health")

	// Inicio del servidor
	log.Fatal(http.ListenAndServe(port, mux))
}
