package main

import "fmt"
import "log"
import (
	"net/http"
	"testing"
)

func TestRun(tst *testing.T) {
	log.Println(" test case starting.....")
	fmt.Println("case OK.")
	mycaserun()
}

func badHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not a regular name or password", http.StatusBadRequest)
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not a regular name or password", http.StatusOK)
}

func headerHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(0)
	for name := range r.Form {
		w.Header().Set(name, r.FormValue(name))
	}
}
