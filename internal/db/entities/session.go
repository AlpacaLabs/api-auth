package entities

import authV1 "github.com/AlpacaLabs/protorepo-auth-go/alpacalabs/auth/v1"

type Session struct {
	ID string
	// TODO use byte slice?
	Secret    string
	AccountID string
}

func (s Session) ToProtobuf() *authV1.Session {
	return &authV1.Session{
		Id:        s.ID,
		Secret:    s.Secret,
		AccountId: s.AccountID,
	}
}

func SessionFromProtobuf(s *authV1.Session) Session {
	return Session{
		ID:        s.Id,
		Secret:    s.Secret,
		AccountID: s.AccountId,
	}
}
