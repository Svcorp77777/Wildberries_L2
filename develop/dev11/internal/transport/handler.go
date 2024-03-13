package transport

import (
	"log"
	"net/http"
	"time"
)

type Handler interface {
	Init(router *http.ServeMux)
}

func LoggingRequest(whiteListHosts []string, h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s %s", time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
	})
}