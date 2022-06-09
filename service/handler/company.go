package handler

import (
	"activeReconBot/dao"
	"activeReconBot/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		log.Printf("[ERROR] %s %s %s %d %s %s", r.RemoteAddr, r.RequestURI, r.Method, r.ContentLength, r.Header.Get("User-Agent"), err.Error())
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	result, err := dao.CreateCompany(r)
	if err != nil {
		log.Printf("[ERROR] %s %s %s %d %s %s", r.RemoteAddr, r.RequestURI, r.Method, r.ContentLength, r.Header.Get("User-Agent"), err.Error())
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, result)
}

func UpdateCompany(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		log.Printf("[ERROR] %s %s %s %d %s %s", r.RemoteAddr, r.RequestURI, r.Method, r.ContentLength, r.Header.Get("User-Agent"), err.Error())
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	companyID := vars["companyID"]
	result, err := dao.UpdateCompany(companyID, r)
	if err != nil {
		log.Printf("[ERROR] %s %s %s %d %s %s", r.RemoteAddr, r.RequestURI, r.Method, r.ContentLength, r.Header.Get("User-Agent"), err.Error())
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, result)
}

func DeleteCompany(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		log.Printf("[ERROR] %s %s %s %d %s %s", r.RemoteAddr, r.RequestURI, r.Method, r.ContentLength, r.Header.Get("User-Agent"), err.Error())
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	companyID := vars["companyID"]
	result, err := dao.DeleteCompany(companyID)
	if err != nil {
		log.Printf("[ERROR] %s %s %s %d %s %s", r.RemoteAddr, r.RequestURI, r.Method, r.ContentLength, r.Header.Get("User-Agent"), err.Error())
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, result)
}

func GetCompanies(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		log.Printf("[ERROR] %s %s %s %d %s %s", r.RemoteAddr, r.RequestURI, r.Method, r.ContentLength, r.Header.Get("User-Agent"), err.Error())
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	result, err := dao.GetCompanies()
	if err != nil {
		log.Printf("[ERROR] %s %s %s %d %s %s", r.RemoteAddr, r.RequestURI, r.Method, r.ContentLength, r.Header.Get("User-Agent"), err.Error())
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, result)
}

func GetCompany(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		log.Printf("[ERROR] %s %s %s %d %s %s", r.RemoteAddr, r.RequestURI, r.Method, r.ContentLength, r.Header.Get("User-Agent"), err.Error())
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	companyID := vars["companyID"]
	result, err := dao.GetCompany(companyID)
	if err != nil {
		log.Printf("[ERROR] %s %s %s %d %s %s", r.RemoteAddr, r.RequestURI, r.Method, r.ContentLength, r.Header.Get("User-Agent"), err.Error())
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, result)
}
