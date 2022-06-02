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

func GetSubdomainsForDomain(domainID string) ([]models.Subdomain, error) {
	var result []models.Subdomain

	_, err := GetDomain(domainID)
	if err != nil {
		//TODO
		return nil, err
	}

	err = mgm.Coll(&models.Subdomain{}).SimpleFind(&result, bson.M{"domain_id": domainID})
	if err != nil {
		//TODO
		return nil, err
	}
	return result, nil
}

func GetSubdomainForDomain(subdomainID string, domainID string) (models.Subdomain, error) {
	var result models.Subdomain

	_, err := GetDomain(domainID)
	if err != nil {
		//TODO
		return result, err
	}

	err = mgm.Coll(&models.Subdomain{}).FindByID(subdomainID, &result)
	if err != nil {
		//TODO
		return result, err
	}
	return result, nil
}

func GetSubdomains() ([]models.Subdomain, error) {
	var result []models.Subdomain
	err := mgm.Coll(&models.Subdomain{}).SimpleFind(&result, bson.M{})
	if err != nil {
		//TODO
		return nil, err
	}
	return result, nil
}

func GetSubdomain(subdomainID string) (models.Subdomain, error) {
	var result models.Subdomain
	err := mgm.Coll(&models.Subdomain{}).FindByID(subdomainID, &result)
	if err != nil {
		//TODO
		return result, err
	}
	return result, nil
}

func CreateSubdomainForDomain(r *http.Request, domainID string) (models.Subdomain, error) {
	decoder := json.NewDecoder(r.Body)
	var subdomain models.Subdomain

	domain, err := GetDomain(domainID)
	if err != nil {
		//TODO
		return subdomain, err
	}

	err = decoder.Decode(&subdomain)
	subdomain.CreatedAt = time.Now()
	subdomain.UpdatedAt = time.Now()
	subdomain.DomainID = domainID
	if err != nil {
		//TODO
		return subdomain, err
	}

	err = mgm.Coll(&domain).Create(&domain)
	if err != nil {
		//TODO
		return subdomain, err
	}

	//Add subdomain to the domain  array
	domain.Subdomains = append(domain.Subdomains, subdomain)
	err = mgm.Coll(&models.Domain{}).Update(&domain)

	//TODO: Add domains to the company array

	return subdomain, nil
}

func CreateSubdomain(r *http.Request) (models.Subdomain, error) {
	decoder := json.NewDecoder(r.Body)
	var subdomain models.Subdomain

	err := decoder.Decode(&subdomain)
	subdomain.CreatedAt = time.Now()
	subdomain.UpdatedAt = time.Now()
	if err != nil {
		//TODO
		return subdomain, err
	}

	err = mgm.Coll(&subdomain).Create(&subdomain)
	if err != nil {
		//TODO
		return subdomain, err
	}

	//Add domain to the company array
	if subdomain.DomainID != "" {
		domain, err := GetDomain(subdomain.DomainID)
		if err != nil {
			//TODO
			return subdomain, err
		}
		domain.Subdomains = append(domain.Subdomains, subdomain)
		err = mgm.Coll(&models.Domain{}).Update(&domain)
		if err != nil {
			//TODO
			return subdomain, err
		}
	}

	return subdomain, nil
}

func UpdateSubdomainForDomain(subdomainID, domainID string, r *http.Request) (models.Subdomain, error) {
	var result = models.Subdomain{}

	domain, err := GetDomain(domainID)
	if err != nil {
		//TODO
		return result, err
	}

	err = mgm.Coll(&models.Subdomain{}).FindByID(subdomainID, &result)
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

	err = mgm.Coll(&models.Subdomain{}).Update(&result)
	if err != nil {
		//TODO
		return result, err
	}

	//Update subdomain in the subdomains array
	var subdomains []models.Subdomain
	for _, subdomain := range domain.Subdomains {
		if subdomain.ID.Hex() == result.ID.Hex() {
			subdomain = result
		}
		subdomains = append(subdomains, subdomain)
	}

	domain.Subdomains = subdomains
	err = mgm.Coll(&models.Domain{}).Update(&domain)
	if err != nil {
		//TODO
		return result, err
	}

	return result, nil
}

func UpdateSubdomain(subdomainID string, r *http.Request) (models.Subdomain, error) {
	var result = models.Subdomain{}

	err := mgm.Coll(&models.Subdomain{}).FindByID(subdomainID, &result)
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

	err = mgm.Coll(&models.Subdomain{}).Update(&result)
	if err != nil {
		//TODO
		return result, err
	}

	//Update domain in the company array
	if result.DomainID != "" {
		domain, err := GetDomain(result.DomainID)
		if err != nil {
			//TODO
			return result, err
		}
		var subdomains []models.Subdomain
		exists := false
		for _, subdomain := range domain.Subdomains {
			if subdomain.ID.Hex() == result.ID.Hex() {
				subdomain = result
				exists = true
			}
			subdomains = append(subdomains, subdomain)
		}

		if !exists {
			subdomains = append(subdomains, result)
		}
		domain.Subdomains = subdomains
		err = mgm.Coll(&models.Domain{}).Update(&domain)
		if err != nil {
			//TODO
			return result, err
		}
	}

	return result, nil
}

func DeleteSubdomainForDomain(domainID, subdomainID string) (map[string]bool, error) {
	var resp = map[string]bool{}
	domain, err := GetDomain(domainID)
	if err != nil {
		//TODO
		resp["success"] = false
		return resp, err
	}

	subdomain := &models.Subdomain{}

	err = mgm.Coll(&models.Subdomain{}).FindByID(subdomainID, subdomain)
	if err != nil {
		//TODO
		return resp, err
	}

	err = mgm.Coll(&models.Subdomain{}).Delete(subdomain)
	if err != nil {
		//TODO
		return resp, err
	}

	//Delete domain from company array
	var subdomains []models.Subdomain
	for idx, subdomainItem := range domain.Subdomains {
		if subdomainItem.ID.Hex() == subdomainID {
			subdomains = utils.RemoveFromSubDomainSlice(domain.Subdomains, idx)
			break
		}
	}

	domain.Subdomains = subdomains
	err = mgm.Coll(&models.Domain{}).Update(&domain)
	if err != nil {
		//TODO
		resp["success"] = false
		return resp, err
	}

	resp["success"] = true
	return resp, nil
}

func DeleteSubdomain(subdomainID string) (map[string]bool, error) {
	var resp = map[string]bool{}

	subdomain := &models.Subdomain{}

	err := mgm.Coll(&models.Subdomain{}).FindByID(subdomainID, subdomain)
	if err != nil {
		//TODO
		return resp, err
	}

	err = mgm.Coll(&models.Subdomain{}).Delete(subdomain)
	if err != nil {
		//TODO
		return resp, err
	}

	//Delete domain from company array
	if subdomain.DomainID != "" {
		domain, err := GetDomain(subdomain.DomainID)
		if err != nil {
			//TODO
			resp["success"] = false
			return resp, err
		}

		var subdomains []models.Subdomain
		for idx, subdomainItem := range domain.Subdomains {
			if subdomainItem.ID.Hex() == subdomainID {
				subdomains = utils.RemoveFromSubDomainSlice(domain.Subdomains, idx)
				break
			}
		}

		domain.Subdomains = subdomains
		err = mgm.Coll(&models.Domain{}).Update(&domain)
		if err != nil {
			//TODO
			resp["success"] = false
			return resp, err
		}
	}

	resp["success"] = true
	return resp, nil
}
