package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GhostsBR/API_Desafio_Go/controller"
)

func main() {
	api_routes() // Need be the last function
}

func api_routes() {
	http.HandleFunc("/api/v1/templates", func (w http.ResponseWriter, r *http.Request) {
		db := database.Database {
			Url: "mongodb+srv://server:jKfERWF0CFlLf8GY@cluster0.agvxp.mongodb.net/myFirstDatabase?retryWrites=true&w=majority",
		}
		result := db.GetTemplates()
		c, err := json.Marshal(result)
		if err != nil {http.Error(w, "Error: cannot convert data into json", 500)}
		w.Write(c)
	})
	fmt.Println("Servidor iniciado!")
	http.ListenAndServe(":5000", nil)
}