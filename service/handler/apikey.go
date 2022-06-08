package handler

import (
	"activeReconBot/dao"
	"activeReconBot/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateApiKey(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	result, err := dao.CreateApiKey(r)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, result)
}

func GetApiKey(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	apiKeyID := vars["apiKeyID"]
	result, err := dao.GetApiKeyByID(apiKeyID)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, result)
}

func GetApiKeys(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	result, err := dao.GetApiKeys()
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, result)
}

func UpdateApiKey(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	apiKeyID := vars["apiKeyID"]

	res, err := dao.GetApiKeyByID(apiKeyID)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	result, err := dao.UpdateApiKey(r, res)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, result)
}

func DeleteApiKey(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	apiKeyID := vars["apiKeyID"]

	res, err := dao.GetApiKeyByID(apiKeyID)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	result, err := dao.DeleteApiKey(res)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, result)
}
