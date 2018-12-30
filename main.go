package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type logRecorder struct {
	http.ResponseWriter

	status int
}

func (lr *logRecorder) WriteHeader(code int) {
	lr.status = code
	lr.ResponseWriter.WriteHeader(code)
}

// PanicHandler attempts to panic but the RecoveryHandler catches the
// panic and returns an Internal Server Error.
func PanicHandler(w http.ResponseWriter, r *http.Request) {
	panic("PANIC")
}

// PostHandler returns Status No Content.
func PostHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

// HealthzHandler returns Status OK.
func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// CatchAllHandler is a catch all route that will return Status Not
// Found for any routes that are not registered in the Router.
func CatchAllHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lr := logRecorder{w, 200}

		start := time.Now()
		next.ServeHTTP(&lr, r)
		elapsed := float64(time.Since(start)) / float64(time.Millisecond)
		log.Printf("\033[32m%s\033[0m %v [%v] %v ms -- %s", r.Method, lr.status, r.URL, elapsed, r.RemoteAddr)
	})
}

func main() {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.Use(handlers.CORS())
	r.Use(handlers.RecoveryHandler())

	r.HandleFunc("/panic", PanicHandler).Methods("GET", "OPTIONS")
	r.HandleFunc("/post", PostHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/healthz", HealthzHandler).Methods("GET", "OPTIONS")
	r.PathPrefix("/").Handler(http.HandlerFunc(CatchAllHandler))

	log.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
