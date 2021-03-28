package user

import (
	"context"
	"github.com/iamsayantan/konference"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	users konference.UserRepository
}

func NewUserService(repo konference.UserRepository) konference.UserService {
	return &userService{users: repo}
}

func (us *userService) CreateUser(ctx context.Context, email, firstName, lastName, plaintextPassword string) error {
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

func (us *userService) Authenticate(ctx context.Context, email, plaintextPassword string) (error, *konference.User) {
	user, err := us.users.FindByEmail(ctx, email)
	if err != nil {
		return err, nil
	}

	if user == nil {
		return konference.ErrInvalidEmailAddress, nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(plaintextPassword))
	if err != nil {
		return konference.ErrInvalidPassword, nil
	}

	return nil, user
}
