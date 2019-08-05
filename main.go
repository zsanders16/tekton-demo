package main

import (
	"html/template"
	"log"
	"net/http"
)

type Message struct {
	Motd string `json:"M"`
}

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/health", Health)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomePage(w http.ResponseWriter, r *http.Request) {

	HomePageVars := Message{
		Motd: "Hello Tekton",
	}

	t, err := template.ParseFiles("homepage.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, HomePageVars)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ok"))
}
