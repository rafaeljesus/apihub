package auth_new

import (
	"fmt"
	"strings"
	"time"

	"code.google.com/p/go.crypto/bcrypt"
	"github.com/backstage/backstage/account_new"
	"github.com/backstage/backstage/errors"
	. "github.com/backstage/backstage/log"
	"github.com/backstage/backstage/util"
)

const (
	EXPIRES_IN_SECONDS  = 24 * 3600
	EXPIRES_TOKEN_CACHE = 10 // time in seconds to remove from expire time.
	TOKEN_TYPE          = "Token"
)

type Authenticatable interface {
	Authenticate(email, password string) (*account_new.User, bool)
	CreateUserToken(*account_new.User) (*account_new.TokenInfo, error)
	UserFromToken(token string) (*account_new.User, error)
	RevokeUserToken(token string) error
}

type auth struct {
	store func() (account_new.Storable, error)
}

func NewAuth(store func() (account_new.Storable, error)) *auth {
	return &auth{store: store}
}

func (a *auth) Authenticate(email, password string) (*account_new.User, bool) {
	// FIXME
	store, err := a.store()
	if err != nil {
		Logger.Warn(err.Error())
		return nil, false
	}
	defer store.Close()

	user, err := store.FindUserByEmail(email)
	if err != nil {
		Logger.Info("Failed trying to find the user '%s' to log in. Original Error: '%s'.", email, err.Error())
		return nil, false
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		Logger.Info("User '%s' is trying to log in with invalid password.", email)
		return nil, false
	}

	return &user, true
}

func (a *auth) CreateUserToken(user *account_new.User) (*account_new.TokenInfo, error) {
	store, err := a.store()
	if err != nil {
		Logger.Warn(err.Error())
		return nil, err
	}
	defer store.Close()

	api := account_new.TokenInfo{
		CreatedAt: time.Now().In(time.UTC).Format("2006-01-02T15:04:05Z07:00"),
		Expires:   EXPIRES_IN_SECONDS,
		Type:      TOKEN_TYPE,
		Token:     util.GenerateRandomStr(32),
		User:      user,
	}

	err = store.CreateToken(api)
	return &api, err
}

func (a *auth) UserFromToken(token string) (*account_new.User, error) {
	h := strings.Split(token, " ")
	if len(h) == 2 {
		apiToken := account_new.TokenInfo{Type: h[0], Token: h[1]}

		if apiToken.Type == TOKEN_TYPE {
			var user account_new.User

			store, err := a.store()
			if err != nil {
				Logger.Warn(err.Error())
				return nil, err
			}
			defer store.Close()

			err = store.DecodeToken(apiToken.Token, &user)
			if err != nil {
				return nil, err
			}
			if user.Email == "" {
				return nil, errors.ErrTokenNotFound
			}

			return &user, nil
		}
	}

	return nil, errors.ErrInvalidTokenFormat
}

func (a *auth) RevokeUserToken(token string) error {
	user, err := a.UserFromToken(token)

	if err == nil { //&& user.Email != "" {
		store, err := a.store()
		if err != nil {
			Logger.Warn(err.Error())
			return err
		}
		defer store.Close()

		key := fmt.Sprintf("%s: %s", TOKEN_TYPE, user.Email)
		err = store.DeleteToken(key)
		if err == nil {
			h := strings.Split(token, " ")
			err = store.DeleteToken(h[1])
		}
	}

	return err
}
