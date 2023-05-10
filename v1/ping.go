package handler

import (
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte("{Pong}"))
}
