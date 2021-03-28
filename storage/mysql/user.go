package mysql

import (
	"context"
	"errors"
	"fmt"

	"github.com/iamsayantan/konference"
	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) konference.UserRepository {
	return &userRepo{db: db}
}

type userRepo struct {
	db *gorm.DB
}

func (u *userRepo) Store(ctx context.Context, user *konference.User) error {
	if user.ID != 0 {
		return errors.New("invalid user details")
	}
	return u.db.WithContext(ctx).Create(&user).Error
}

func (u *userRepo) Update(ctx context.Context, user *konference.User) error {
	return nil
}

func (u *userRepo) FindById(ctx context.Context, uid uint) (*konference.User, error) {
	return u.findByField(ctx, "id", uid)
}

func (u *userRepo) FindByEmail(ctx context.Context, email string) (*konference.User, error) {
	return u.findByField(ctx, "email", email)
}

func (u *userRepo) findByField(ctx context.Context, fieldName string, fieldValue interface{}) (*konference.User, error) {
	var user konference.User
	err := u.db.WithContext(ctx).Where(fmt.Sprintf("%s = ?", fieldName), fieldValue).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}
