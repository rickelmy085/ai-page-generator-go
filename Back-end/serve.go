package main

import "net/http"

func main() {
	http.HandleFunc("/AI_qroq", nil)
	http.ListenAndServe(":8080", nil)

}

func HandleAIQroq(w http.ResponseWriter, r *http.Request) {

}
