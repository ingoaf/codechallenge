package database_test

import (
	"testing"

	"github.com/Clarilab/codechallenge/backend/challenge2/database"
	"github.com/Clarilab/codechallenge/backend/challenge2/models"
	"github.com/stretchr/testify/assert"
)

func Test_Repository_Get(t *testing.T) {
	repo := database.NewRepository()
	generalMsg := "wrong search result"
	tests := []struct {
		msg string
		// Input
		search models.SearchRequest
		// Output
		registrationNumbers []string
	}{
		{
			msg:                 "results of functions Get and GetByName do not match",
			search:              models.SearchRequest{CompanyName: "finux"},
			registrationNumbers: []string{"HRB 17675", "HRB 17686"},
		},
		{
			msg:                 generalMsg,
			search:              models.SearchRequest{City: "Kassel"},
			registrationNumbers: []string{"HRB 17675", "HRB 17234"},
		},
		{
			msg:                 generalMsg,
			search:              models.SearchRequest{City: "Kassel", CompanyName: "fino"},
			registrationNumbers: []string{"HRB 17234"},
		},
	}
	for _, test := range tests {
		result := repo.Get(test.search)
		numbers := getRegistrationNumbers(t, result)
		assert.Equal(t, test.registrationNumbers, numbers, test.msg)
	}
}

func getRegistrationNumbers(t *testing.T, companies []*models.Company) []string {
	t.Helper()
	registrationNumbers := make([]string, len(companies))
	for i := range companies {
		registrationNumbers[i] = companies[i].RegistrationNumber
	}
	return registrationNumbers
}
