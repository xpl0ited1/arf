package handler

import "net/http"

func CreateDomain(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, "Testing Api")
}

func GetDomain(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, "Testing Api")
}

func UpdateDomain(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, "Testing Api")
}

func DeleteDomain(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, "Testing Api")
}
