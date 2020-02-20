package service

import (
	"github.com/ProxeusApp/proxeus-core/sys"
	"github.com/ProxeusApp/proxeus-core/sys/model"
)

type (
	UserService interface {
		GetUser(auth model.Auth) (*model.User, error)
	}
	defaultUserService struct {
		*baseService
	}
)

func NewUserService(system *sys.System) *defaultUserService {
	return &defaultUserService{&baseService{system: system}}
}

func (me *defaultUserService) GetUser(auth model.Auth) (*model.User, error) {
	return me.userDB().Get(auth, auth.UserID())
}