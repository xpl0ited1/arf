package dao

import (
	"activeReconBot/models"
	"activeReconBot/utils"
	"encoding/json"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

func CreateUser(r *http.Request) (models.User, error) {
	decoder := json.NewDecoder(r.Body)
	var user models.User
	err := decoder.Decode(&user)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	if err != nil {
		//TODO
		return user, err
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return user, err
	}

	user.Password = hashedPassword

	err = mgm.Coll(&user).Create(&user)
	if err != nil {
		//TODO
		return user, err
	}

	return user, nil
}

func GetUserByID(id string) (models.User, error) {
	var result models.User
	err := mgm.Coll(&models.User{}).FindByID(id, &result)
	result.Password = ""
	if err != nil {
		//TODO
		return result, err
	}
	return result, nil
}

func GetUserByUsername(username string) (models.User, error) {
	var result models.User
	err := mgm.Coll(&models.User{}).First(bson.M{"username": username}, &result)
	if err != nil {
		//TODO
		return result, err
	}
	return result, nil
}

func GetUsers() ([]models.User, error) {
	var result []models.User
	err := mgm.Coll(&models.User{}).SimpleFind(&result, bson.M{})
	if err != nil {
		//TODO
		return nil, err
	}

	for _, user := range result {
		user.Password = ""
	}

	return result, nil
}

func UpdateUser(r *http.Request, user models.User) (models.User, error) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		//TODO
		return user, err
	}

	user.UpdatedAt = time.Now()
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return user, err
	}

	user.Password = hashedPassword

	err = mgm.Coll(&models.User{}).Update(&user)
	if err != nil {
		//TODO
		return user, err
	}

	return user, nil
}

func DeleteUser(user models.User) (map[string]bool, error) {
	var resp = map[string]bool{}

	err := mgm.Coll(&models.User{}).Delete(&user)
	if err != nil {
		//TODO
		return resp, err
	}
	resp["success"] = true
	return resp, nil
}
