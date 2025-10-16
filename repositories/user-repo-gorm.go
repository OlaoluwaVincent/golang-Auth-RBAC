package repositories

import (
	"errors"
	"go/auth/entities"
	"go/auth/interfaces"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID           uint `gorm:"primary_key"`
	Name         string
	Email        string `gorm:"unique_index"`
	PasswordHash string
	Role         string
}

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) Create(u *entities.User) error {
	user := User{
		Name:         u.Name,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		Role:         u.Role,
	}
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}
	u.ID = user.ID
	return nil
}

func (r *GormUserRepository) FindByEmail(email string) (*entities.User, error) {
	var user User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &entities.User{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		Role:         user.Role,
	}, nil
}

func (r *GormUserRepository) FindByID(id uint) (*entities.User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &entities.User{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		Role:         user.Role,
	}, nil
}

func (r *GormUserRepository) Update(u *entities.User) error {
	var user User
	if err := r.db.First(&user, u.ID).Error; err != nil {
		return errors.New("not found")
	}
	user.Name = u.Name
	user.Email = u.Email
	user.PasswordHash = u.PasswordHash
	user.Role = u.Role
	return r.db.Save(&user).Error
}
