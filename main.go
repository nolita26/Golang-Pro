package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

type Data struct {
    Message string `json:"message"`
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/api/data", GetData).Methods("GET")
    fs := http.FileServer(http.Dir("frontend"))
    http.Handle("/", fs)
    log.Fatal(http.ListenAndServe(":8080", r))
}

func GetData(w http.ResponseWriter, r *http.Request) {
    data := Data{Message: "Hello from backend!"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}
