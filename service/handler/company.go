package handler

import (
	"activeReconBot/dao"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	result, err := dao.CreateCompany(r)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, result)
}

func UpdateCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyID := vars["company_id"]
	result, err := dao.UpdateCompany(companyID, r)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, result)
}

func DeleteCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyID := vars["company_id"]
	result, err := dao.DeleteCompany(companyID)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, result)
}

func GetCompanies(w http.ResponseWriter, r *http.Request) {
	result, err := dao.GetCompanies()
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, result)
}

func GetCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyID := vars["company_id"]
	result, err := dao.GetCompany(companyID)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, result)
}
