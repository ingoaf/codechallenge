package database_test

import (
	"testing"

	"github.com/Clarilab/codechallenge/backend/challenge2/database"
)

func Test_Repository_GetByName(t *testing.T) {
	repo := database.NewRepository()

	result := repo.GetByName("finux")

	if result == nil {
		t.Error("Expected to find any result")
	}

	if len(result) != 2 {
		t.Error("Expected to find two results")
	}
}
