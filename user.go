package konference

import (
	"context"
	"errors"
	"time"
)

var (
	ErrEmailAlreadyTaken   = errors.New("the email is already taken")
	ErrInvalidEmailAddress = errors.New("invalid email id")
	ErrInvalidPassword     = errors.New("invalid password")
	ErrUserNotFound        = errors.New("user not found")
)

// User type represents the User domain in the system.
type User struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Email     string    `json:"email" gorm:"unique"`
	FirstName string    `json:"firstName" gorm:"size:191"`
	LastName  string    `json:"lastName" gorm:"size:191"`
	Password  string    `json:"-" gorm:"size:255"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// NewUser creates a new user
func NewUser(email, firstName, lastName, password string) *User {
	return &User{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Password:  password,
	}
}

// UserRepository defines the methods to interact with the User storage.
type UserRepository interface {
	Store(ctx context.Context, u *User) error
	Update(ctx context.Context, u *User) error
	FindById(ctx context.Context, id uint) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
}

// UserService defines the domain service methods for User domain.
type UserService interface {
	// CreateUser creates a new user in the system.
	CreateUser(ctx context.Context, email, firstName, lastName, plaintextPassword string) error
	// Authenticate verifies the credentials and returns the associated user if they match.
	Authenticate(ctx context.Context, email, password string) (*User, error)
	// GetUserDetails returns the details of the user
	GetUserDetails(ctx context.Context, userId uint) (*User, error)
}
