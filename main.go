package main

import (
	"encoding/json"
	"fmt"
	"github.com/GhostsBR/API_Desafio_Go/controller"
	"net/http"
	"strconv"
)

func main() {
	api_routes() // Need be the last function
}


func api_routes() {
	db := database.Database {
		Url: "mongodb+srv://system:WV4fNKP2axPC5eUv@cluster0.agvxp.mongodb.net/myFirstDatabase?retryWrites=true&w=majority",
	}
	http.HandleFunc("/api/v1/templates", func (w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			result := db.GetTemplates()
			c, err := json.Marshal(result)
			if err != nil {http.Error(w, "Error: cannot convert data into json", 500); return}
			w.Write(c)
		} else {
			http.Error(w, "Error: Invalid method, please use GET method.", 400)
			return
		}
	})
	http.HandleFunc("/api/v1/template", func (w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			keys, ok := r.URL.Query()["id"]
			if !ok || len(keys[0]) < 1 {http.Error(w, "Error: query id not found!", 400); return}
			id := keys[0]
			intid, errc := strconv.Atoi(id)
			if errc != nil {http.Error(w, "Error: cannot id in int", 500); return}
			result := db.GetTemplate(intid)
			c, err := json.Marshal(result)
			if err != nil {http.Error(w, "Error: cannot convert data into json", 500); return}
			w.Write(c)
		} else if r.Method == http.MethodPost {
			//var t database.Template
			//err := json.NewDecoder(r.Body).Decode(&t)
			//if err != nil {
			//	http.Error(w, err.Error(), http.StatusBadRequest)
			//	return
			//}


			//w.Header().Set("Content-Type", "application/json")
			//body, err := ioutil.ReadAll(r.Body)
			//var data interface{}
			//if err != nil {http.Error(w, "Error: Cannot read body", 400)}
			////json.Unmarshal(body, &data)
			//err = json.NewDecoder(r.Body).Decode(&data)
			//if err != nil {http.Error(w, "Error: Cannot read body", 400)}
			////w.Write(body)
		//	if db.InsertData() {
		//		fmt.Fprintln(w, "Sucess: template created with sucess!")
		//	} else {
		//		http.Error(w, "Error: template cannot be added in database!", 400)
		//	}
		}
	})
	fmt.Println("Servidor iniciado!")
	http.ListenAndServe(":5000", nil)
}