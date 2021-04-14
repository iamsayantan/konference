package room

import (
	"context"
	"github.com/iamsayantan/konference"
)

type roomServiceImpl struct {
	rooms konference.RoomRepository
	us    konference.UserService
}

func (rs *roomServiceImpl) Create(ctx context.Context, ownerId uint) (*konference.Room, error) {
	panic("implement me")
}

func (rs *roomServiceImpl) GetDetails(ctx context.Context, invitationCode string) (*konference.Room, error) {
	panic("implement me")
}

func (rs *roomServiceImpl) Join(ctx context.Context, roomId uint, joiningUserId uint) error {
	panic("implement me")
}

func (rs *roomServiceImpl) Leave(ctx context.Context, roomId uint, leavingUserId uint) error {
	panic("implement me")
}

func (rs *roomServiceImpl) IsResiding(ctx context.Context, roomId uint, userId uint) bool {
	panic("implement me")
}

func NewRoomService(r konference.RoomRepository, us konference.UserService) konference.RoomService {
	return &roomServiceImpl{
		rooms: r,
		us:    us,
	}
}
