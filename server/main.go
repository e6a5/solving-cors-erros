package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	addr := ":8000"
	router := mux.NewRouter()
	router.HandleFunc("/users", handleList)
	log.Println("listen on", fmt.Sprintf("http://localhost%s/", addr))
	log.Fatal(http.ListenAndServe(addr, router))
}

type ResponseData struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

func handleList(w http.ResponseWriter, res *http.Request) {
	log.Println("Request Method ", res.Method)
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	if res.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	test1, err := res.Cookie("test1")
	if err != nil || test1 == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp := []ResponseData{
		{
			Name: "Hiep Tran",
			Age:  int32(26),
		},
	}
	bytes, _ := json.Marshal(resp)
	io.WriteString(w, string(bytes))
}
