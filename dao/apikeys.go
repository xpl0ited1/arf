package dao

import (
	"activeReconBot/models"
	"encoding/json"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

func GetApiKeys() ([]models.ApiKey, error) {
	var result []models.ApiKey
	err := mgm.Coll(&models.ApiKey{}).SimpleFind(&result, bson.M{})
	if err != nil {
		//TODO
		return nil, err
	}
	return result, nil
}

func GetApiKeyByID(id string) (models.ApiKey, error) {
	var result models.ApiKey
	err := mgm.Coll(&models.ApiKey{}).FindByID(id, &result)
	if err != nil {
		//TODO
		return result, err
	}
	return result, nil
}

func GetApiKeyByKey(key string) (models.ApiKey, error) {
	var result models.ApiKey
	err := mgm.Coll(&models.ApiKey{}).SimpleFind(&result, bson.M{"apikey": key})
	if err != nil {
		//TODO
		return result, err
	}
	return result, nil
}

func CreateApiKey(r *http.Request) (models.ApiKey, error) {
	decoder := json.NewDecoder(r.Body)
	var apiKey models.ApiKey
	err := decoder.Decode(&apiKey)
	apiKey.CreatedAt = time.Now()
	apiKey.UpdatedAt = time.Now()
	if err != nil {
		//TODO
		return apiKey, err
	}

	err = mgm.Coll(&apiKey).Create(&apiKey)
	if err != nil {
		//TODO
		return apiKey, err
	}

	return apiKey, nil
}

func UpdateApiKey(r *http.Request, apiKey models.ApiKey) (models.ApiKey, error) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&apiKey)
	if err != nil {
		//TODO
		return apiKey, err
	}

	apiKey.UpdatedAt = time.Now()

	err = mgm.Coll(&models.ApiKey{}).Update(&apiKey)
	if err != nil {
		//TODO
		return apiKey, err
	}

	return apiKey, nil
}

func DeleteApiKey(apiKey models.ApiKey) (map[string]bool, error) {
	var resp = map[string]bool{}

	err := mgm.Coll(&models.ApiKey{}).Delete(&apiKey)
	if err != nil {
		//TODO
		return resp, err
	}
	resp["success"] = true
	return resp, nil
}
