package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/NYTimes/gziphandler"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	//"gopkg.in/mgo.v2"
)

type Welcome struct {
	Name string
	Time string
}

func home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("frontend/src/index.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, nil)
}

func main() {
	// welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}
	// templates := template.Must(template.ParseFiles("frontend/src/index.html"))
	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/api/v1/signup", signup).Methods("POST")
	r.HandleFunc("/api/v1/login", login).Methods("GET")

	fmt.Println("Serving on PORT 8080")
	// fmt.Println(http.ListenAndServe(":8080", nil))

	server := &http.Server{
		Addr:    ":8080",
		Handler: cors.Default().Handler(gziphandler.GzipHandler(noCacheMW(r))),
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
