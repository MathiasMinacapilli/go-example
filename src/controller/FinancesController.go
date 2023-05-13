
package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

import "../services/FinancesService"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hola, mundo!")
	})

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		// llama a la función que obtiene los datos de la base de datos
		data, err := FinancesService.getDataFromDatabase()
		if err != nil {
			http.Error(w, "Error al obtener los datos", http.StatusInternalServerError)
			return
		}

		// Establecer el encabezado de la respuesta como JSON
		w.Header().Set("Content-Type", "application/json")

		// Codificar los datos como JSON y escribirlos en la respuesta
		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			http.Error(w, "Error al codificar los datos", http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":8080", nil)
}
