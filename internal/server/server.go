package server

import (
	"net/http"
)

func StartServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/getData", getStats)
	mux.Handle("/", http.FileServer(http.Dir("/home/kirill/serverVizzorStatic")))

	http.ListenAndServe(":8080", mux)
}
