package api

import (
	"github.com/salmaan72/latitude-assignment/internal/auth"
	"github.com/salmaan72/latitude-assignment/internal/user"
	"github.com/salmaan72/latitude-assignment/internal/user/ledger"
)

type API struct {
	UserService   *user.Service
	AuthService   *auth.Service
	LedgerService *ledger.Service
}

func New(userService *user.Service, authService *auth.Service, ledgerService *ledger.Service) *API {
	return &API{
		UserService:   userService,
		AuthService:   authService,
		LedgerService: ledgerService,
	}
}
