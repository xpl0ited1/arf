package handler

import (
	"activeReconBot/models"
	"activeReconBot/utils"
	"encoding/json"
	"net/http"
)

func DummyLogin(w http.ResponseWriter, r *http.Request) {
	var user models.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	var response = map[string]string{}
	if user.Username == "ACAB" && user.Password == "473f5985cea3ac848546e9ef178e9ae1018b4bef57bdc6be408609a4532cb179" {
		token, err := utils.CreateTokenForUser(user.Username, user.Username)
		if err != nil {
			RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		response["token"] = token
		RespondJSON(w, http.StatusOK, response)
	} else {
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
}
