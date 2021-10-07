package konference

import (
	"context"
)

// Room is the place where a call takes place.
type Room struct {
	ID      int64 `json:"id"`
	OwnerID uint  `json:"-"`
	Owner   *User `json:"createdBy" gorm:"foreignKey:OwnerID"`
}

func NewRoom(id int64, owner *User) *Room {
	room := &Room{
		ID:      id,
		OwnerID: owner.ID,
		Owner:   owner,
	}

	return room
}

// RoomRepository defines the methods to interact with the room storage.
type RoomRepository interface {
	Store(ctx context.Context, room *Room) error
	FindById(ctx context.Context, id int64) (*Room, error)
}

// RoomService defines the application methods exposed by the rooms domain.
type RoomService interface {
	// Create creates a new room. Creating a room does not add the creator as
	// a member by default.
	Create(ctx context.Context, ownerId uint) (*Room, error)

	// GetDetails returns room details by the id.
	GetDetails(ctx context.Context, id int64) (*Room, error)
}
