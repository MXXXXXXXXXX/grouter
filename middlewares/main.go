package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/xujiajun/gorouter"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

//serve static resources
func ServeFiles(w http.ResponseWriter, r *http.Request) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dir := wd + "/middlewares/files"
	http.StripPrefix("/files/", http.FileServer(http.Dir(dir))).ServeHTTP(w, r)
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}

//https://upgear.io/blog/golang-tip-wrapping-http-response-writer-for-middleware/
func withStatusRecord(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rec := statusRecorder{w, http.StatusOK}
		next.ServeHTTP(&rec, r)
		log.Printf("response status: %v\n", rec.status)
	}
}

func notFoundFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Not found page !")
}

func withLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Logged connection from %s", r.RemoteAddr)
		next.ServeHTTP(w, r)
	}
}

func withTracing(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Tracing request for %s", r.RequestURI)
		next.ServeHTTP(w, r)
	}
}

func main() {
	mux := gorouter.New()
	mux.NotFoundFunc(notFoundFunc)
	mux.Use(withLogging, withTracing, withStatusRecord)
	mux.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	//defined prefix
	mux2 := mux.Group("/files")
	//http://127.0.0.1:8181/files/demo.txt
	//will match
	mux2.GET("/{filename:[0-9a-zA-Z_.]+}", func(w http.ResponseWriter, r *http.Request) {
		ServeFiles(w, r)
	})

	log.Fatal(http.ListenAndServe(":34401", mux))
}
