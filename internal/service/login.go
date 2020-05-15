package service

import (
	"context"
	"strings"

	authV1 "github.com/AlpacaLabs/protorepo-auth-go/alpacalabs/auth/v1"
	"github.com/badoux/checkmail"
	"github.com/dgrijalva/jwt-go"
	"github.com/ttacon/libphonenumber"
)

// AlpacaGroups are an extension of jwt.StandardClaims, with roles.
type AlpacaClaims struct {
	// Groups a list of the current user's groups
	//Groups []string `json:"groups,omitempty"`
	jwt.StandardClaims
}

func (s Service) Login(ctx context.Context, request *authV1.LoginRequest) (*authV1.LoginResponse, error) {
	// Validate email w/ checkmail
	// Retrieve account entity by username, email, or phone number
	// Call matchesHash function
	// Persist login attempt?
	// Check if MFA is enabled
	// Create AlpacaClaims
	// Generate per-token secret
	// Persist secret
	// Create JWT
	// In the HTTP controller, return Set-Cookie header w/ JWT

	return nil, nil
}

func isEmailAddress(s string) bool {
	return strings.Contains(s, "@") && checkmail.ValidateFormat(s) == nil
}

func isPhoneNumber(s string) bool {
	_, err := libphonenumber.Parse(s, "US")
	return err == nil
}
