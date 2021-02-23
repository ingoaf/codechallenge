package models

type SearchResponse struct {
	TotalHits int       `json:"totalHits,omitempty" redis:"totalHits"`
	Companies []Company `json:"companies,omitempty" redis:"companies"`
}

type Company struct {
	Name               string   `json:"name,omitempty" redis:"name,omitempty"`
	RegistrationNumber string   `json:"registrationNumber,omitempty" redis:"registrationNumber,omitempty"`
	Address            Address  `json:"address,omitempty" redis:"address,omitempty"`
	Provider           Provider `json:"provider,omitempty" redis:"provider,omitempty"`
	AlreadyOnboarded   bool     `json:"alreadyOnboarded,omitempty" redis:"alreadyOnboarded,omitempty"`
}

type Address struct {
	Country     string `json:"country,omitempty" redis:"country,omitempty"`
	CountryCode string `json:"countryCode,omitempty" redis:"countryCode,omitempty"`
	StreetLine  string `json:"streetLine,omitempty" redis:"streetLine,omitempty"`
	Locality    string `json:"locality,omitempty" redis:"locality,omitempty"`
}

type Provider struct {
	Name       string `json:"name" redis:"name,omitempty"`
	ProviderID string `json:"providerID" redis:"providerID,omitempty"`
	IDKey      string `json:"idKey" redis:"idKey,omitempty"`
}
