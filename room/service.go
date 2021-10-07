package room

import (
	"context"

	"github.com/iamsayantan/konference"
	"github.com/iamsayantan/konference/util"
)

type roomServiceImpl struct {
	rooms konference.RoomRepository
	us    konference.UserService
}

func (rs *roomServiceImpl) Create(ctx context.Context, ownerId uint) (*konference.Room, error) {
	user, err := rs.us.GetUserDetails(ctx, ownerId)
	if err != nil {
		return nil, err
	}

	id := util.GenerateNewId()
	room := konference.NewRoom(id, user)

	err = rs.rooms.Store(ctx, room)
	if err != nil {
		return nil, err
	}

	return rs.rooms.FindById(ctx, id)
}

func (rs *roomServiceImpl) GetDetails(ctx context.Context, id int64) (*konference.Room, error) {
	return rs.rooms.FindById(ctx, id)
}

func NewRoomService(r konference.RoomRepository, us konference.UserService) konference.RoomService {
	return &roomServiceImpl{
		rooms: r,
		us:    us,
	}
}
