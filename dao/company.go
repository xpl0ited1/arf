package dao

import (
	"activeReconBot/models"
	"encoding/json"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

func GetCompanies() ([]models.Company, error) {
	var result []models.Company
	err := mgm.Coll(&models.Company{}).SimpleFind(&result, bson.M{})
	if err != nil {
		//TODO
		return nil, err
	}
	return result, nil
}

func GetCompany(companyID string) (models.Company, error) {
	var result models.Company
	err := mgm.Coll(&models.Company{}).FindByID(companyID, &result)
	if err != nil {
		//TODO
		return result, err
	}
	return result, nil
}

func CreateCompany(r *http.Request) (models.Company, error) {
	decoder := json.NewDecoder(r.Body)
	var company models.Company
	err := decoder.Decode(&company)
	company.CreatedAt = time.Now()
	company.UpdatedAt = time.Now()
	if err != nil {
		//TODO
		return company, err
	}

	err = mgm.Coll(&company).Create(&company)
	if err != nil {
		//TODO
		return company, err
	}

	return company, nil
}

func UpdateCompany(companyID string, r *http.Request) (models.Company, error) {
	var result = models.Company{}
	err := mgm.Coll(&models.Company{}).FindByID(companyID, &result)
	if err != nil {
		//TODO
		return result, err
	}

	createdAt := result.CreatedAt

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&result)
	if err != nil {
		//TODO
		return result, err
	}

	result.CreatedAt = createdAt
	result.UpdatedAt = time.Now()

	err = mgm.Coll(&models.Company{}).Update(&result)
	if err != nil {
		//TODO
		return result, err
	}

	return result, nil
}

func DeleteCompany(companyID string) (map[string]bool, error) {
	company := &models.Company{}
	var resp = map[string]bool{}
	err := mgm.Coll(&models.Company{}).FindByID(companyID, company)
	if err != nil {
		//TODO
		return resp, err
	}

	err = mgm.Coll(&models.Company{}).Delete(company)
	if err != nil {
		//TODO
		return resp, err
	}
	resp["success"] = true
	return resp, nil
}
