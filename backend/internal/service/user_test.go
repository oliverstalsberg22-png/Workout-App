package service

import (
	"github.com/Oliverstalsy/egolifter/internal/domain"
	"os"
	"testing"
)

func TestCreateToken(t *testing.T) {
	os.Setenv("JWTKEY", "testing_key")
	user := &domain.User{
		ID: "123yomama",
	}
	token, err := CreateToken(user)
	if err != nil {
		t.Fatalf("Error, %v", err)
	}
	if token == "" {
		t.Fatal("Did not create token")
	}
}
