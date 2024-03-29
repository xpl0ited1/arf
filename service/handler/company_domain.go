package handler

import (
	"activeReconBot/dao"
	"activeReconBot/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func CreateDomainForCompany(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		log.Printf("[ERROR] %s %s %s %d %s %s", r.RemoteAddr, r.RequestURI, r.Method, r.ContentLength, r.Header.Get("User-Agent"), err.Error())
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	companyID := vars["companyID"]
	result, err := dao.CreateDomainForCompany(r, companyID)
	if err != nil {
		log.Printf("[ERROR] %s %s %s %d %s %s", r.RemoteAddr, r.RequestURI, r.Method, r.ContentLength, r.Header.Get("User-Agent"), err.Error())
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, result)
}

func CreateDomain(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		log.Printf("[ERROR] %s %s %s %d %s %s", r.RemoteAddr, r.RequestURI, r.Method, r.ContentLength, r.Header.Get("User-Agent"), err.Error())
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	result, err := dao.CreateDomain(r)
	if err != nil {
		log.Printf("[ERROR] %s %s %s %d %s %s", r.RemoteAddr, r.RequestURI, r.Method, r.ContentLength, r.Header.Get("User-Agent"), err.Error())
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, result)
}

func GetDomainForCompany(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		log.Printf("[ERROR] %s %s %s %d %s %s", r.RemoteAddr, r.RequestURI, r.Method, r.ContentLength, r.Header.Get("User-Agent"), err.Error())
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	companyID := vars["companyID"]
	domainID := vars["domainID"]
	result, err := dao.GetDomainForCompany(companyID, domainID)
	if err != nil {
		log.Printf("[ERROR] %s %s %s %d %s %s", r.RemoteAddr, r.RequestURI, r.Method, r.ContentLength, r.Header.Get("User-Agent"), err.Error())
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, result)
}

func GetDomains(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-Api-Key")
	if apiKey != "" {
		res, err := dao.GetApiKeyByKey(apiKey)
		if err != nil {
			log.Printf("[ERROR] %s %s %s %d %s %s", r.RemoteAddr, r.RequestURI, r.Method, r.ContentLength, r.Header.Get("User-Agent"), err.Error())
			RespondError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		if !res.Enabled {
			RespondError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		if res.ApiKey == "" {
			RespondError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
	} else {
		_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
		if err != nil {
			log.Printf("[ERROR] %s %s %s %d %s %s", r.RemoteAddr, r.RequestURI, r.Method, r.ContentLength, r.Header.Get("User-Agent"), err.Error())
			RespondError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
	}

	result, err := dao.GetDomains()
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, result)
}

func GetDomain(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	domainID := vars["domainID"]
	result, err := dao.GetDomain(domainID)
	if err != nil {
		log.Printf("[ERROR] %s %s %s %d %s %s", r.RemoteAddr, r.RequestURI, r.Method, r.ContentLength, r.Header.Get("User-Agent"), err.Error())
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, result)
}

func GetDomainsForCompany(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		log.Printf("[ERROR] %s %s %s %d %s %s", r.RemoteAddr, r.RequestURI, r.Method, r.ContentLength, r.Header.Get("User-Agent"), err.Error())
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	companyID := vars["companyID"]
	result, err := dao.GetDomainsForCompany(companyID)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, result)
}

func UpdateDomainForCompany(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		log.Printf("[ERROR] %s %s %s %d %s %s", r.RemoteAddr, r.RequestURI, r.Method, r.ContentLength, r.Header.Get("User-Agent"), err.Error())
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	companyID := vars["companyID"]
	domainID := vars["domainID"]
	result, err := dao.UpdateDomainForCompany(domainID, companyID, r)
	if err != nil {
		log.Printf("[ERROR] %s %s %s %d %s %s", r.RemoteAddr, r.RequestURI, r.Method, r.ContentLength, r.Header.Get("User-Agent"), err.Error())
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, result)
}

func UpdateDomain(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		log.Printf("[ERROR] %s %s %s %d %s %s", r.RemoteAddr, r.RequestURI, r.Method, r.ContentLength, r.Header.Get("User-Agent"), err.Error())
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	domainID := vars["domainID"]
	result, err := dao.UpdateDomain(domainID, r)
	if err != nil {
		log.Printf("[ERROR] %s %s %s %d %s %s", r.RemoteAddr, r.RequestURI, r.Method, r.ContentLength, r.Header.Get("User-Agent"), err.Error())
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, result)
}

func DeleteDomain(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		log.Printf("[ERROR] %s %s %s %d %s %s", r.RemoteAddr, r.RequestURI, r.Method, r.ContentLength, r.Header.Get("User-Agent"), err.Error())
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	domainID := vars["domainID"]

	result, err := dao.DeleteDomain(domainID)
	if err != nil {
		log.Printf("[ERROR] %s %s %s %d %s %s", r.RemoteAddr, r.RequestURI, r.Method, r.ContentLength, r.Header.Get("User-Agent"), err.Error())
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, result)
}

func DeleteDomainForCompany(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		log.Printf("[ERROR] %s %s %s %d %s %s", r.RemoteAddr, r.RequestURI, r.Method, r.ContentLength, r.Header.Get("User-Agent"), err.Error())
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	companyID := vars["companyID"]
	domainID := vars["domainID"]

	result, err := dao.DeleteDomainForCompany(domainID, companyID)
	if err != nil {
		log.Printf("[ERROR] %s %s %s %d %s %s", r.RemoteAddr, r.RequestURI, r.Method, r.ContentLength, r.Header.Get("User-Agent"), err.Error())
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, result)
}
