package database

import (
	"strings"

	"github.com/Clarilab/codechallenge/backend/challenge4/models"
)

// Repository loads data from the database
type Repository interface {
	Get(request *models.SearchRequest) ([]models.Company, error)
}

type repository struct {
	Companies []models.Company
}

// NewRepository is a repository constructor
func NewRepository() Repository {
	return &repository{Companies: getAll()}
}

// Get returns all companies which match criteria from the search request
func (repo *repository) Get(request *models.SearchRequest) ([]models.Company, error) {
	companies := repo.Companies
	companiesMatchingSearch := []models.Company{}
	for _, com := range companies {
		if companyMatchesSearch(&com, *request) {
			companiesMatchingSearch = append(companiesMatchingSearch, com)
		}
	}
	return companiesMatchingSearch, nil
}

func companyMatchesSearch(company *models.Company, request models.SearchRequest) bool {
	match := false
	companyName := strings.ToLower(company.Name)
	request.CompanyName = strings.ToLower(request.CompanyName)
	if request.CompanyName != "" && !strings.Contains(companyName, request.CompanyName) {
		return match
	}
	if request.Country != "" && request.Country != company.Address.Country {
		return match
	}
	if request.City != "" && request.City != company.Address.Locality {
		return match
	}

	// only if unknown parameters are empty the search is succesful, otherwise we assume
	// no entry matches the given criteria
	return request.Owner == "" && request.Region == "" && request.ZipCode == ""
}

func getAll() []models.Company {
	return []models.Company{
		{
			Name:               "Finux GmbH",
			RegistrationNumber: "HRB 17675",
			Address: models.Address{
				Country:     "Deutschland",
				CountryCode: "DE",
				Locality:    "Kassel",
				StreetLine:  "Universitätsplatz 12",
			},
		},
		{
			Name:               "Finux GmbH",
			RegistrationNumber: "HRB 17686",
			Address: models.Address{
				Country:     "Deutschland",
				CountryCode: "DE",
				Locality:    "Osnabrück",
				StreetLine:  "Bahnhofstraße 23",
			},
		},
		{
			Name:               "Linux GmbH",
			RegistrationNumber: "HRA 634675",
			Address: models.Address{
				Country:     "Österreich",
				CountryCode: "AU",
				Locality:    "Wien",
				StreetLine:  "Opernplatz 42",
			},
		},
		{
			Name:               "Finex GmbH & Co. KG",
			RegistrationNumber: "HRB 345342",
			Address: models.Address{
				Country:     "Deutschland",
				CountryCode: "DE",
				Locality:    "Frankfurt am Main",
				StreetLine:  "Lange Str. 1234454",
			},
		},
		{
			Name:               "fino Create GmbH",
			RegistrationNumber: "HRB 17234",
			Address: models.Address{
				Country:     "Deutschland",
				CountryCode: "DE",
				Locality:    "Kassel",
				StreetLine:  "Universitätsplatz 12",
			},
		},
	}
}
