package fiano

import (
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler is the main Handler for a path in the URL
type Handler struct {
	*mux.Router
}

// ServeStatic serves static resources
func (hd Handler) ServeStatic(path string) {
	hd.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		cnt, err := ioutil.ReadFile(r.URL.Path)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
		}
		w.Write(cnt)
	})
}

func NewHandler(r *mux.Router) *Handler {
	return &Handler{r}
}
