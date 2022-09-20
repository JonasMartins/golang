package main

import (
	"fmt"
	"net/http"
	"regexp"
	"sync"
)

var (
	listUserRe   = regexp.MustCompile(`^\/api/users[\/]*$`)
	getUserRe    = regexp.MustCompile(`^\/api/users\/(\d+)$`)
	createUserRe = regexp.MustCompile(`^\/api/users[\/]*$`)
)

// user represents our REST resource
type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// our in-memory datastore
// rememeber to guard map access with a mutex for concurrent access
type datastore struct {
	m map[string]user
	*sync.RWMutex
}
type handler struct {
	store *datastore
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet && listUserRe.MatchString(r.URL.Path):
		h.List(w, r)
		return
	case r.Method == http.MethodGet && getUserRe.MatchString(r.URL.Path):
		h.Get(w, r)
		return
	case r.Method == http.MethodPost && createUserRe.MatchString(r.URL.Path):
		h.Create(w, r)
		return
	default:
		notFound(w, r)
		return
	}
}

func main() {

	mux := http.NewServeMux()
	userH := &handler{
		store: &datastore{
			m: map[string]user{
				"1": user{ID: "1", Name: "bob"},
			},
			RWMutex: &sync.RWMutex{},
		},
	}

	port := 4000

	mux.Handle("/api/users", userH)
	mux.Handle("/api/users/", userH)

	fmt.Printf("Server running at port %d\n", port)
	http.ListenAndServe("localhost:4000", mux)
}
