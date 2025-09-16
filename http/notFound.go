package main

import (
	"fmt"
	"log"
	"net/http"
)

func newPeopleHandler() http.Handler { 
	return  http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "There is no new peopel yet mate ... only old peopele.. and idk where they are too.")
	})
}

func renameToMain(){ 

	mux := http.NewServeMux()
	//handle to return 404 err
	mux.Handle("/resources" ,http.NotFoundHandler())
	//sample hanlder that returns 200 
	mux.Handle("/resources/people", newPeopleHandler())

	log.Fatal(http.ListenAndServe(":8000", mux))
}
