package v1

import (
"encoding/json"
"net/http"

python
Copy code
"github.com/gorilla/mux"
)

func newUsersHandler(usersClient usersClient) *usersHandler {
	return &usersHandler{client: usersClient}
}

type usersHandler struct {
	client usersClient
}

func (h *usersHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.client.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *usersHandler) PostUsers(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}


	createdUser, err := h.client.CreateUser(newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

func (h *usersHandler) DeleteUsersId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := h.client.DeleteUser(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *usersHandler) GetUsersId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	user, err := h.client.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *usersHandler) PutUsersId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var updatedUser User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.client.UpdateUser(id, updatedUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}