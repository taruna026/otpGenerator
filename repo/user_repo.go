package repo

import (
	"github.com/jinzhu/gorm"
	"sinarmas/models"
)

type userRepo struct {
	db *gorm.DB
}

//go:generate mockgen -source=./user_repo.go -package=mocks -destination=./mocks/user_repo.go
type IUserRepo interface {
	Create(user *models.User) error
	GetByUserAndRequestId(userId, requestId string) (*models.User, error)
	Save(user *models.User) error
}

func NewUserRepo(db *gorm.DB) IUserRepo {
	return &userRepo{db: db}
}

func (u userRepo) Create(user *models.User) error {
	return u.db.Create(user).Error
}

func (u userRepo) GetByUserAndRequestId(userId, requestId string) (*models.User, error) {
	var userInfo models.User

	err := u.db.Where("user_id=? and request_id=? and is_validated=false", userId, requestId).First(&userInfo)
	if err.Error != nil {
		return nil, err.Error
	}

	return &userInfo, nil
}

func (u userRepo) Save(user *models.User) error {
	return u.db.Save(user).Error
}
