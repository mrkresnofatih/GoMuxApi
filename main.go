package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	mux "github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello World")

	r := mux.NewRouter()
	s := r.PathPrefix("/home").Subrouter()
	s.Use(loggingMIddleware)
	s.HandleFunc("/test/{id}", HomeHandler).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":3000", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	h := r.Header.Get("testheader")
	p := mux.Vars(r)["id"]
	log.Println("testheader: " + h)
	log.Println("id: " + p)
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, "Test Header is "+h)
}

func loggingMIddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("req uri:" + r.RequestURI)

		ht := r.Header.Get("sampleId")
		if ht == "1234" {
			err := errors.New("SampleId header is incorrect")
			if err != nil {
				panic("sdfjlksdfj")
			}
			w.WriteHeader(http.StatusBadRequest)
			f, e := json.Marshal(map[string]string{
				"test_a": "haha",
				"test_b": "hbhb",
			})
			if e != nil {
				log.Println(e)
			}

			fmt.Fprint(w, string(f))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
