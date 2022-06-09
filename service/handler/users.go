package handler

import (
	"activeReconBot/dao"
	"activeReconBot/models"
	"activeReconBot/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
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
	if user.Username == "ACAB" && user.Password == "" {
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

func CreateUser(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	result, err := dao.CreateUser(r)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, result)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	userID := vars["userID"]

	res, err := dao.GetUserByID(userID)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	result, err := dao.DeleteUser(res)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, result)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	userID := vars["userID"]

	res, err := dao.GetUserByID(userID)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	result, err := dao.UpdateUser(r, res)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, result)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	vars := mux.Vars(r)
	userID := vars["userID"]
	result, err := dao.GetUserByID(userID)
	if err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, result)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetUserIDFromToken(r.Header.Get("Authorization"))
	if err != nil {
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	result, err := dao.GetUsers()
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, result)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var response = map[string]string{}

	users, err := dao.GetUsers()
	if err != nil {
		log.Println(err)
	}

	if len(users) == 0 {
		token, _ := utils.CreateTokenForUser("dummy", "dummy")
		response["token"] = token
		RespondJSON(w, http.StatusOK, response)
		return
	}

	var user models.User
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&user)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	dbUser, err := dao.GetUserByUsername(user.Username)
	if err != nil {
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if user.Username == "" {
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if !utils.CheckPasswordHash(user.Password, dbUser.Password) {
		RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	token, err := utils.CreateTokenForUser(user.ID.Hex(), user.Username)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response["token"] = token
	RespondJSON(w, http.StatusOK, response)
}
