package fiano

import (
	"log"
	"net/http"
)

type Runner struct {
	hdl http.Handler
}

// AddHandler adds a new handler to the list
func (rn *Runner) AddHandler(hdl Handler) {
	old := rn.hdl
	rn.hdl = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		old.ServeHTTP(w, r)
		hdl.ServeHTTP(w, r)
	})
}

// Run runs the server
func (rn *Runner) Run(addr string) {
	http.Handle("/", rn.hdl)
	log.Fatal(http.ListenAndServe(addr, nil))
}
