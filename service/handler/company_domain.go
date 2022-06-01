package handler

import (
	"activeReconBot/dao"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateDomainForCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyID := vars["companyID"]
	result, err := dao.CreateDomainForCompany(r, companyID)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, result)
}

func CreateDomain(w http.ResponseWriter, r *http.Request) {
	result, err := dao.CreateDomain(r)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, result)
}

func GetDomainForCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyID := vars["companyID"]
	domainID := vars["domainID"]
	result, err := dao.GetDomainForCompany(companyID, domainID)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, result)
}

func GetDomains(w http.ResponseWriter, r *http.Request) {
	result, err := dao.GetDomains()
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, result)
}

func GetDomain(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	domainID := vars["domainID"]
	result, err := dao.GetDomain(domainID)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, result)
}

func GetDomainsForCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyID := vars["companyID"]
	result, err := dao.GetDomainsForCompany(companyID)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, result)
}

func UpdateDomainForCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyID := vars["companyID"]
	domainID := vars["domainID"]
	result, err := dao.UpdateDomainForCompany(domainID, companyID, r)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, result)
}

func UpdateDomain(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	domainID := vars["domainID"]
	result, err := dao.UpdateDomain(domainID, r)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, result)
}

func DeleteDomain(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	domainID := vars["domainID"]

	result, err := dao.DeleteDomain(domainID)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, result)
}

func DeleteDomainForCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyID := vars["companyID"]
	domainID := vars["domainID"]

	result, err := dao.DeleteDomainForCompany(domainID, companyID)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, result)
}
