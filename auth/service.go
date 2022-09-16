package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Service interface {
	FetchAuth(AuthDetails) (*Auth, error)
	DeleteAuth(AuthDetails) (bool, error)
	CreateAuth(uint64) (*Auth, error)
	Login(AuthDetails) (string, error)
	AuthMiddleware() gin.HandlerFunc
	FindAuthUser(userID int) bool
	DeleteAuthUser(userID int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FetchAuth(authD AuthDetails) (*Auth, error) {

	auths, err := s.repository.FindAuth(&authD)
	if err != nil {
		return nil, err
	}

	if auths == nil {
		return nil, err
	} else {
		au, err := s.repository.FetchAuth(auths)
		if err != nil {
			return nil, err
		}

		return au, nil
	}

}

func (s *service) DeleteAuth(authD AuthDetails) (bool, error) {

	auths, err := s.repository.FindAuth(&authD)
	if err != nil {
		return false, err
	}

	if auths == nil {
		return false, nil
	} else {
		err = s.repository.DeleteAuth(auths)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func (s *service) CreateAuth(userID uint64) (*Auth, error) {
	au, err := s.repository.CreateAuth(userID)
	if err != nil {
		return nil, err
	}
	return au, nil
}

func (s *service) Login(authD AuthDetails) (string, error) {
	token, err := CreateToken(authD)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *service) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "You need to be authorized to access this route")
			c.Abort()
			return
		}
		c.Next()
	}
}

func (s *service) FindAuthUser(userID int) bool {
	err := s.repository.FindAuthUser(userID)
	if err {
		return true
	} else {
		return false
	}
}

func (s *service) DeleteAuthUser(userID int) error {
	err := s.repository.DeleteAuthUser(userID)
	if err != nil {
		return err
	}
	return nil
}
