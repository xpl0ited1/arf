package handler

import (
	"activeReconBot/dao"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateSubdomain(w http.ResponseWriter, r *http.Request) {
	result, err := dao.CreateSubdomain(r)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, result)
}

func GetSubdomain(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subdomainID := vars["subdomainID"]
	result, err := dao.GetSubdomain(subdomainID)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, result)
}

func GetSubdomains(w http.ResponseWriter, r *http.Request) {
	result, err := dao.GetSubdomains()
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, result)
}

func UpdateSubdomain(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subdomainID := vars["subdomainID"]
	result, err := dao.UpdateSubdomain(subdomainID, r)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, result)
}

func DeleteSubdomain(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subdomainID := vars["subdomainID"]

	result, err := dao.DeleteSubdomain(subdomainID)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, result)
}

func CreateSubdomainForDomain(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	domainID := vars["domainID"]
	result, err := dao.CreateSubdomainForDomain(r, domainID)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, result)
}

func GetSubdomainForDomain(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	domainID := vars["domainID"]
	subdomainID := vars["subdomainID"]
	result, err := dao.GetSubdomainForDomain(subdomainID, domainID)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, result)
}

func GetSubdomainsForDomain(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	domainID := vars["domainID"]
	result, err := dao.GetSubdomainsForDomain(domainID)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, result)
}

func UpdateSubdomainForDomain(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subdomainID := vars["subdomainID"]
	domainID := vars["domainID"]
	result, err := dao.UpdateSubdomainForDomain(subdomainID, domainID, r)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, result)
}

func DeleteSubdomainForDomain(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subdomainID := vars["subdomainID"]
	domainID := vars["domainID"]
	result, err := dao.DeleteSubdomainForDomain(domainID, subdomainID)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, result)
}
