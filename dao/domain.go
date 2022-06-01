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

func GetDomainsForCompany(companyID string) ([]models.Domain, error) {
	var result []models.Domain

	_, err := GetCompany(companyID)
	if err != nil {
		//TODO
		return nil, err
	}

	err = mgm.Coll(&models.Domain{}).SimpleFind(&result, bson.M{"company_id": companyID})
	if err != nil {
		//TODO
		return nil, err
	}
	return result, nil
}

func GetDomainForCompany(companyID string, domainID string) (models.Domain, error) {
	var result models.Domain

	_, err := GetCompany(companyID)
	if err != nil {
		//TODO
		return result, err
	}

	err = mgm.Coll(&models.Domain{}).FindByID(domainID, &result)
	if err != nil {
		//TODO
		return result, err
	}
	return result, nil
}

func GetDomains() ([]models.Domain, error) {
	var result []models.Domain
	err := mgm.Coll(&models.Domain{}).SimpleFind(&result, bson.M{})
	if err != nil {
		//TODO
		return nil, err
	}
	return result, nil
}

func GetDomain(domainID string) (models.Domain, error) {
	var result models.Domain
	err := mgm.Coll(&models.Domain{}).FindByID(domainID, &result)
	if err != nil {
		//TODO
		return result, err
	}
	return result, nil
}

func CreateDomainForCompany(r *http.Request, companyID string) (models.Domain, error) {
	decoder := json.NewDecoder(r.Body)
	var domain models.Domain

	company, err := GetCompany(companyID)
	if err != nil {
		//TODO
		return domain, err
	}

	err = decoder.Decode(&domain)
	domain.CreatedAt = time.Now()
	domain.UpdatedAt = time.Now()
	domain.CompanyID = companyID
	if err != nil {
		//TODO
		return domain, err
	}

	err = mgm.Coll(&domain).Create(&domain)
	if err != nil {
		//TODO
		return domain, err
	}

	//Add domain to the company array
	company.Domains = append(company.Domains, domain)
	err = mgm.Coll(&models.Company{}).Update(&company)

	return domain, nil
}

func CreateDomain(r *http.Request) (models.Domain, error) {
	decoder := json.NewDecoder(r.Body)
	var domain models.Domain

	err := decoder.Decode(&domain)
	domain.CreatedAt = time.Now()
	domain.UpdatedAt = time.Now()
	if err != nil {
		//TODO
		return domain, err
	}

	err = mgm.Coll(&domain).Create(&domain)
	if err != nil {
		//TODO
		return domain, err
	}

	//Add domain to the company array
	if domain.CompanyID != "" {
		company, err := GetCompany(domain.CompanyID)
		if err != nil {
			//TODO
			return domain, err
		}
		company.Domains = append(company.Domains, domain)
		err = mgm.Coll(&models.Company{}).Update(&company)
		if err != nil {
			//TODO
			return domain, err
		}
	}

	return domain, nil
}

func UpdateDomainForCompany(domainID, companyID string, r *http.Request) (models.Domain, error) {
	var result = models.Domain{}

	company, err := GetCompany(companyID)
	if err != nil {
		//TODO
		return result, err
	}

	err = mgm.Coll(&models.Domain{}).FindByID(domainID, &result)
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

	err = mgm.Coll(&models.Domain{}).Update(&result)
	if err != nil {
		//TODO
		return result, err
	}

	//Update domain in the company array
	var domains []models.Domain
	for _, domain := range company.Domains {
		if domain.ID.Hex() == result.ID.Hex() {
			domain = result
		}
		domains = append(domains, domain)
	}

	company.Domains = domains
	err = mgm.Coll(&models.Company{}).Update(&company)
	if err != nil {
		//TODO
		return result, err
	}

	return result, nil
}

func UpdateDomain(domainID string, r *http.Request) (models.Domain, error) {
	var result = models.Domain{}

	err := mgm.Coll(&models.Domain{}).FindByID(domainID, &result)
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

	err = mgm.Coll(&models.Domain{}).Update(&result)
	if err != nil {
		//TODO
		return result, err
	}

	//Update domain in the company array
	if result.CompanyID != "" {
		company, err := GetCompany(result.CompanyID)
		if err != nil {
			//TODO
			return result, err
		}
		var domains []models.Domain
		exists := false
		for _, domain := range company.Domains {
			if domain.ID.Hex() == result.ID.Hex() {
				domain = result
				exists = true
			}
			domains = append(domains, domain)
		}

		if !exists {
			domains = append(domains, result)
		}
		company.Domains = domains
		err = mgm.Coll(&models.Company{}).Update(&company)
		if err != nil {
			//TODO
			return result, err
		}
	}

	return result, nil
}

func DeleteDomainForCompany(domainID, companyID string) (map[string]bool, error) {
	var resp = map[string]bool{}
	company, err := GetCompany(companyID)
	if err != nil {
		//TODO
		resp["success"] = false
		return resp, err
	}

	domain := &models.Domain{}

	err = mgm.Coll(&models.Domain{}).FindByID(domainID, domain)
	if err != nil {
		//TODO
		return resp, err
	}

	err = mgm.Coll(&models.Domain{}).Delete(domain)
	if err != nil {
		//TODO
		return resp, err
	}

	//Delete domain from company array
	var domains []models.Domain
	for idx, domainItem := range company.Domains {
		if domainItem.ID.Hex() == domainID {
			domains = utils.RemoveFromDomainSlice(company.Domains, idx)
			break
		}
	}

	company.Domains = domains
	err = mgm.Coll(&models.Company{}).Update(&company)
	if err != nil {
		//TODO
		resp["success"] = false
		return resp, err
	}

	resp["success"] = true
	return resp, nil
}

func DeleteDomain(domainID string) (map[string]bool, error) {
	var resp = map[string]bool{}

	domain := &models.Domain{}

	err := mgm.Coll(&models.Domain{}).FindByID(domainID, domain)
	if err != nil {
		//TODO
		return resp, err
	}

	err = mgm.Coll(&models.Domain{}).Delete(domain)
	if err != nil {
		//TODO
		return resp, err
	}

	//Delete domain from company array
	if domain.CompanyID != "" {
		company, err := GetCompany(domain.CompanyID)
		if err != nil {
			//TODO
			resp["success"] = false
			return resp, err
		}

		var domains []models.Domain
		for idx, domainItem := range company.Domains {
			if domainItem.ID.Hex() == domainID {
				domains = utils.RemoveFromDomainSlice(company.Domains, idx)
				break
			}
		}

		company.Domains = domains
		err = mgm.Coll(&models.Company{}).Update(&company)
		if err != nil {
			//TODO
			resp["success"] = false
			return resp, err
		}
	}

	resp["success"] = true
	return resp, nil
}
