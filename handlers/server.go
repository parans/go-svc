package handlers

import (
	"net/http"

	gmux "github.com/gorilla/mux"
	"github.td.teradata.com/appcenter/backup-restore/internal"
)

//Server ...
type Server struct {
	ds internal.DataStore
}

//NewServer ...
func NewServer(dtstr internal.DataStore) *Server {
	return &Server{
		ds: dtstr,
	}
}

// Mux is a multiplexer which defines middleware and all API paths
func (s *Server) Mux() http.Handler {
	r := gmux.NewRouter()

	v1 := r.PathPrefix("/v1").Subrouter()
	v1.StrictSlash(true)
	v1.Methods("GET").Path("/backup").Handler(HTTPHandler{handlerFunc: s.Charts})
	v1.Methods("POST").Path("/backup").Handler(HTTPHandler{handlerFunc: s.UploadChart})

	// Add more paths here.
	return r
}

// Error is a wrapper inerface for HTTP error
type Error interface {
	error
	Status() int
}

// HTTPError implements error and Error interface
type HTTPError struct {
	code int
	err  error
}

// Error implements error interface
func (h HTTPError) Error() string {
	return h.err.Error()
}

// Status implements Error interface
func (h HTTPError) Status() int {
	return h.code
}

// HTTPHandler ... is a wrapper function that returns error from http request
type HTTPHandler struct {
	handlerFunc func(w http.ResponseWriter, r *http.Request) error
}

// ServeHttp handles http request
func (h HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.handlerFunc(w, r)
	if err != nil {
		switch e := err.(type) {
		case HTTPError:
			http.Error(w, e.Error(), e.Status())
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError)+":"+err.Error(), http.StatusInternalServerError)
		}

	}
}
