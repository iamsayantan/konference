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
	user, err := rs.us.GetUserDetails(ctx, ownerId)
	if err != nil {
		return nil, err
	}

	room := konference.NewRoom(user)

	err = rs.rooms.Store(ctx, room)
	if err != nil {
		return nil, err
	}

	return rs.rooms.FindByInviteCode(ctx, room.InviteCode)
}

func (rs *roomServiceImpl) GetDetails(ctx context.Context, invitationCode string) (*konference.Room, error) {
	return rs.rooms.FindByInviteCode(ctx, invitationCode)
}

func (rs *roomServiceImpl) Join(ctx context.Context, roomId uint, joiningUserId uint) error {
	user, err := rs.us.GetUserDetails(ctx, joiningUserId)
	if err != nil {
		return err
	}

	room, err := rs.rooms.FindById(ctx, roomId)
	if err != nil {
		return err
	}

	err = room.AddMember(user)
	if err != nil {
		return err
	}

	return nil
}

func (rs *roomServiceImpl) Leave(ctx context.Context, roomId uint, leavingUserId uint) error {
	room, err := rs.rooms.FindById(ctx, roomId)
	if err != nil {
		return err
	}

	err = room.RemoveMember(leavingUserId)
	if err != nil {
		return err
	}

	return nil
}

func (rs *roomServiceImpl) IsResiding(ctx context.Context, roomId uint, userId uint) bool {
	room, err := rs.rooms.FindById(ctx, roomId)
	if err != nil {
		return false
	}

	member := room.GetMember(userId)
	return member != nil
}

func NewRoomService(r konference.RoomRepository, us konference.UserService) konference.RoomService {
	return &roomServiceImpl{
		rooms: r,
		us:    us,
	}
}
