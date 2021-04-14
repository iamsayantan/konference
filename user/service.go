package user

import (
	"context"
	"github.com/iamsayantan/konference"
	"golang.org/x/crypto/bcrypt"
)

type userServiceImpl struct {
	users konference.UserRepository
}

func NewUserService(repo konference.UserRepository) konference.UserService {
	return &userServiceImpl{users: repo}
}

func (us *userServiceImpl) CreateUser(ctx context.Context, email, firstName, lastName, plaintextPassword string) error {
	existingUser, err := us.users.FindByEmail(ctx, email)
	if err != nil {
		return err
	}

	if existingUser != nil {
		return konference.ErrEmailAlreadyTaken
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), bcrypt.DefaultCost)
	if err != nil {
		return konference.ErrInternalServerError
	}

	user := konference.NewUser(email, firstName, lastName, string(hashedPassword))
	return us.users.Store(ctx, user)
}

func (us *userServiceImpl) Authenticate(ctx context.Context, email, plaintextPassword string) (*konference.User, error) {
	user, err := us.users.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, konference.ErrInvalidEmailAddress
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(plaintextPassword))
	if err != nil {
		return nil, konference.ErrInvalidPassword
	}

	return user, nil
}

func (us *userServiceImpl) GetUserDetails(ctx context.Context, userId uint) (*konference.User, error) {
	user, err := us.users.FindById(ctx, userId)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, konference.ErrUserNotFound
	}

	return user, nil
}
