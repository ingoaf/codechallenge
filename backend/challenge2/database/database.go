package database

import "github.com/Clarilab/codechallenge/backend/challenge2/models"

// Repository loads data from the database
type Repository interface {
	getAll() []models.Company
	GetByName(name string) []models.Company
	Get(request models.SearchRequest) []models.Company
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) GetByName(name string) []models.Company {
	panic("implement me")
}

func (repo *repository) Get(request models.SearchRequest) []models.Company {
	panic("implement me")
}

func (repo *repository) getAll() []models.Company {
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
