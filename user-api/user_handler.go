package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type UserHandler struct {
	store UserStore
}

func NewUserHandler(store UserStore) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

func (uh *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/users")
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/")

	switch r.Method {
	case "GET":
		uh.getHandler(w, r)
	case "POST":
		uh.createHandler(w, r)
	case "DELETE":
		uh.deleteHandler(w, r)
	}
}

func (uh *UserHandler) getHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "" {
		uh.listHandler(w, r)
		return
	}

	userName := r.URL.Path
	user, err := uh.store.Get(userName)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	payload, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(payload)
}

func (uh *UserHandler) listHandler(w http.ResponseWriter, r *http.Request) {
	users, err := uh.store.List()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	payload, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(payload)
}

func (uh *UserHandler) createHandler(w http.ResponseWriter, r *http.Request) {
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	user := &User{}
	err = json.Unmarshal(payload, user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if user.Name == "" {
		http.Error(w, "invalid user: empty username", 400)
		return
	}

	err = uh.store.Create(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(201)
}

func (uh *UserHandler) deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "" {
		http.Error(w, "no user specified", 404)
		return
	}

	userName := r.URL.Path

	err := uh.store.Delete(userName)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(204)
}
