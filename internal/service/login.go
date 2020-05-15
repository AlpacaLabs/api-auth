package service

import (
	"context"
	"strings"

	accountV1 "github.com/AlpacaLabs/protorepo-account-go/alpacalabs/account/v1"
	authV1 "github.com/AlpacaLabs/protorepo-auth-go/alpacalabs/auth/v1"
	passwordV1 "github.com/AlpacaLabs/protorepo-password-go/alpacalabs/password/v1"
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
	// Get account
	accountClient := accountV1.NewAccountServiceClient(s.accountConn)
	accountRes, err := accountClient.GetAccount(ctx, getAccountRequest(request.AccountIdentifier))
	if err != nil {
		return nil, err
	}

	accountID := accountRes.Account.Id

	// Check password
	passwordClient := passwordV1.NewPasswordServiceClient(s.passwordConn)
	if _, err := passwordClient.CheckPassword(ctx, &passwordV1.CheckPasswordRequest{
		AccountId: accountID,
		Password:  request.Password,
	}); err != nil {
		return nil, err
	}

	// TODO Check if MFA is required
	// TODO persist Session to DB
	// TODO return JWT

	// Create AlpacaClaims
	// Generate per-token secret
	// Create JWT
	var jwtString string

	return &authV1.LoginResponse{
		Jwt: jwtString,
	}, nil
}

// the input string can be a email address, phone number, or username
func getAccountRequest(accountIdentifier string) *accountV1.GetAccountRequest {
	if isEmailAddress(accountIdentifier) {
		return &accountV1.GetAccountRequest{
			AccountIdentifier: &accountV1.GetAccountRequest_EmailAddress{
				EmailAddress: accountIdentifier,
			},
		}
	} else if isPhoneNumber(accountIdentifier) {
		return &accountV1.GetAccountRequest{
			AccountIdentifier: &accountV1.GetAccountRequest_PhoneNumber{
				PhoneNumber: accountIdentifier,
			},
		}
	}

	return &accountV1.GetAccountRequest{
		AccountIdentifier: &accountV1.GetAccountRequest_Username{
			Username: accountIdentifier,
		},
	}
}

func isEmailAddress(s string) bool {
	return strings.Contains(s, "@") && checkmail.ValidateFormat(s) == nil
}

func isPhoneNumber(s string) bool {
	_, err := libphonenumber.Parse(s, "US")
	return err == nil
}
