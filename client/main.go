package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	addr := ":8080"
	http.HandleFunc("/", homePage)
	log.Println("listen on", fmt.Sprintf("http://localhost%s/", addr))
	log.Fatal(http.ListenAndServe(addr, nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	ts.Execute(w, nil)
}
