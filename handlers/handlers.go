package handlers

import (
	"net/http"
	"fmt"
	"encoding/json"
)
import "go-example/model"

func GetUsers(response http.ResponseWriter, request *http.Request) {
	users := []model.User{
		{ID: 1, Name: "Alice", Email: "alice@example.com"},
		{ID: 2, Name: "Bob", Email: "bob@example.com"},
		{ID: 3, Name: "Charlie", Email: "charlie@example.com"},
	}

	// Convertir los usuarios a JSON
	jsonData, err := json.Marshal(users)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	// Configurar la cabecera de la respuesta para indicar que se trata de un JSON
	response.Header().Set("Content-Type", "application/json")

	// Escribir la respuesta
	response.Write(jsonData)
}

func GetUser(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "obtiene un usuario")
}