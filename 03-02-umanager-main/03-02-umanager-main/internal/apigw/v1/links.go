package v1

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func newLinksHandler(linksClient linksClient) *linksHandler {
	return &linksHandler{client: linksClient}
}

type linksHandler struct {
	client linksClient
}

func (h *linksHandler) GetLinks(w http.ResponseWriter, r *http.Request) {
	links, err := h.client.GetLinks()
	if err != nil {
		http.Error(w, "Failed to get links", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(links)
}

func (h *linksHandler) PostLinks(w http.ResponseWriter, r *http.Request) {
	var link Link
	err := json.NewDecoder(r.Body).Decode(&link)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	createdLink, err := h.client.CreateLink(link)
	if err != nil {
		http.Error(w, "Failed to create link", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdLink)
}

func (h *linksHandler) DeleteLinksId(w http.ResponseWriter, r *http.Request, id string) {
	err := h.client.DeleteLink(id)
	if err != nil {
		http.Error(w, "Failed to delete link", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *linksHandler) GetLinksId(w http.ResponseWriter, r *http.Request, id string) {
	link, err := h.client.GetLinkByID(id)
	if err != nil {
		http.Error(w, "Failed to get link", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(link)
}

func (h *linksHandler) PutLinksId(w http.ResponseWriter, r *http.Request, id string) {
	var link Link
	err := json.NewDecoder(r.Body).Decode(&link)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	updatedLink, err := h.client.UpdateLink(id, link)
	if err != nil {
		http.Error(w, "Failed to update link", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedLink)
}

func (h *linksHandler) GetLinksUserUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]
	links, err := h.client.GetLinksByUserID(userID)

	if err != nil {
		http.Error(w, "Failed to get links for user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(links)
}
