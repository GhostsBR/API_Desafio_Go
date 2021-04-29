package main

import (
	"github.com/GhostsBR/API_Desafio_Go/controller"
	"fmt"
	"net/http"
)

func main() {
	//  db := database.Database {
	//  	Url: "mongodb+srv://server:jKfERWF0CFlLf8GY@cluster0.agvxp.mongodb.net/myFirstDatabase?retryWrites=true&w=majority",
	// }
	//db.InitDatabase()
	api_routes() // Need be the last function
}

func api_routes() {
	http.HandleFunc("/api/v1/templates", func (w http.ResponseWriter, r *http.Request) {
		db := database.Database {
			Url: "mongodb+srv://server:jKfERWF0CFlLf8GY@cluster0.agvxp.mongodb.net/myFirstDatabase?retryWrites=true&w=majority",
		}
		fmt.Fprintln(w, db.GetTemplates())
	})


	fmt.Println("Servidor iniciado com sucesso!")
	http.ListenAndServe(":5000", nil)
}