package auth

import (
	"github.com/twinj/uuid"
	"gorm.io/gorm"
)

type Auth struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	UserID   uint64 `gorm:"not null" json:"user_id"`
	AuthUUID string `gorm:"size:255;not null" json:"auth_uuid"`
}

type Repository interface {
	FindAuth(authD *AuthDetails) (*Auth, error)
	FetchAuth(authD *Auth) (*Auth, error)
	DeleteAuth(authD *Auth) error
	CreateAuth(uint64) (*Auth, error)
	FindAuthUser(userID int) bool
	DeleteAuthUser(userID int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAuth(authD *AuthDetails) (*Auth, error) {
	au := &Auth{}
	err := r.db.Where("user_id = ? AND auth_uuid = ?", authD.UserID, authD.AuthUUID).Find(&au).Error
	if err != nil {
		return nil, err
	}
	return au, nil
}

func (r *repository) FetchAuth(authD *Auth) (*Auth, error) {
	err := r.db.Take(&authD).Error
	if err != nil {
		return nil, err
	}
	return authD, nil
}

func (r *repository) DeleteAuth(authD *Auth) error {
	err := r.db.Delete(&authD).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) CreateAuth(userID uint64) (*Auth, error) {
	au := &Auth{}
	au.AuthUUID = uuid.NewV4().String()
	au.UserID = userID
	err := r.db.Create(&au).Error
	if err != nil {
		return nil, err
	}
	return au, nil
}

func (r *repository) FindAuthUser(userID int) bool {
	au := &Auth{}
	err := r.db.Where("user_id = ?", userID).Find(&au).Error
	return err == nil
}

func (r *repository) DeleteAuthUser(userID int) error {
	au := &Auth{}
	err := r.db.Where("user_id = ?", userID).Find(&au).Delete(&au).Error
	if err != nil {
		return err
	}
	return nil
}
