package handler

import (
	"activeReconBot/utils"
	"net/http"
)

func DummyLogin(w http.ResponseWriter, r *http.Request) {
	token, err := utils.CreateTokenForUser("testid", "test")
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	var response = map[string]string{}
	response["token"] = token
	RespondJSON(w, http.StatusOK, response)
}
