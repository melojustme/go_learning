package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

//http简单实现
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("哈哈哈"))

}
