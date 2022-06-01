package handler

import "net/http"

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, "Testing Api")
}

func UpdateCompany(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, "Testing Api")
}

func DeleteCompany(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, "Testing Api")
}

func GetCompanies(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, "Testing Api")
}

func GetCompany(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, "Testing Api")
}
