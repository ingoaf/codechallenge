package models

type SearchRequest struct {
	CompanyName string `json:"companyName,omitempty"`
	Owner       string `json:"owner,omitempty"`
	Country     string `json:"country,omitempty"`
	Region      string `json:"region,omitempty"`
	ZipCode     string `json:"zipCode,omitempty"`
	City        string `json:"city,omitempty"`
	Street      string `json:"street,omitempty"`
}
