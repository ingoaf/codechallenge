package models

type SearchResponse struct {
	TotalHits int       `json:"totalHits,omitempty"`
	Companies []Company `json:"companies,omitempty"`
}

type Company struct {
	Name               string   `json:"name,omitempty"`
	RegistrationNumber string   `json:"registrationNumber,omitempty"`
	Address            Address  `json:"address,omitempty"`
	Provider           Provider `json:"provider,omitempty"`
	AlreadyOnboarded   bool     `json:"alreadyOnboarded,omitempty"`
}

type Address struct {
	Country     string `json:"country,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
	StreetLine  string `json:"streetLine,omitempty"`
	Locality    string `json:"locality,omitempty"`
}

type Provider struct {
	Name       string `json:"name"`
	ProviderID string `json:"providerID"`
	IDKey      string `json:"idKey"`
}
