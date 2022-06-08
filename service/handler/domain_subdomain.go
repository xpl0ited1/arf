package handler

import (
	"activeReconBot/dao"
	"activeReconBot/scanners"
	"activeReconBot/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sync"
)

func CreateSubdomain(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	result, err := dao.CreateSubdomain(r)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, result)
}

func GetSubdomain(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	subdomainID := vars["subdomainID"]
	result, err := dao.GetSubdomain(subdomainID)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, result)
}

func GetSubdomains(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	result, err := dao.GetSubdomains()
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, result)
}

func UpdateSubdomain(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	subdomainID := vars["subdomainID"]
	result, err := dao.UpdateSubdomain(subdomainID, r)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, result)
}

func DeleteSubdomain(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	subdomainID := vars["subdomainID"]

	result, err := dao.DeleteSubdomain(subdomainID)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, result)
}

func CreateSubdomainForDomain(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-Api-Key")
	if apiKey != "" {
		//TODO Implement API KEYS
		if apiKey != "" {
			RespondError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
	} else {
		_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
		if err != nil {
			RespondError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
	}

	var wg sync.WaitGroup

	vars := mux.Vars(r)
	domainID := vars["domainID"]
	result, err := dao.CreateSubdomainForDomain(r, domainID)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		ips := scanners.ResolveIpFromHostnameShodan(result.SubdomainName)
		for _, ip := range ips {
			shodanData := scanners.GetShodanDetailsFromIPAddress(ip.IP)
			dao.UpdateSubdomainNoReq(result.ID.Hex(), shodanData)
		}

		pageTitle, err := scanners.GetPageTitle(result.SubdomainName)
		if err != nil {
			log.Println(err)
		}

		if pageTitle != "" {
			result.HTTPTitle = pageTitle
			result, err = dao.UpdateSubdomainHTTPTitle(result)
			if err != nil {
				log.Println(err)
			}
		}

	}(&wg)

	RespondJSON(w, http.StatusCreated, result)

	wg.Wait()
}

func GetSubdomainForDomain(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	domainID := vars["domainID"]
	subdomainID := vars["subdomainID"]
	result, err := dao.GetSubdomainForDomain(subdomainID, domainID)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, result)
}

func GetSubdomainsForDomain(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	domainID := vars["domainID"]
	result, err := dao.GetSubdomainsForDomain(domainID)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, result)
}

func UpdateSubdomainForDomain(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	subdomainID := vars["subdomainID"]
	domainID := vars["domainID"]
	result, err := dao.UpdateSubdomainForDomain(subdomainID, domainID, r)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, result)
}

func DeleteSubdomainForDomain(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	subdomainID := vars["subdomainID"]
	domainID := vars["domainID"]
	result, err := dao.DeleteSubdomainForDomain(domainID, subdomainID)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, result)
}
