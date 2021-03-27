package konference

import (
	"context"
	"time"
)

// User type represents the User domain in the system.
type User struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Email     string    `json:"email" gorm:"unique"`
	FirstName string    `json:"firstName" gorm:"size:191"`
	LastName  string    `json:"lastName" gorm:"size:191"`
	Password  string    `json:"-" gorm:"size:255"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
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

type UserService interface {
	CreateUser(ctx context.Context, u *User) error
	Authenticate(ctx context.Context, email, password string) (error, *User)
}
