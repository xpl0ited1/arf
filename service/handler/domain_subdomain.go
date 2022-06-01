package handler

import "net/http"

func CreateSubdomain(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, "Testing Api")
}

func GetSubdomain(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, "Testing Api")
}

func UpdateSubdomain(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, "Testing Api")
}

func DeleteSubdomain(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, "Testing Api")
}
