package mysql

import (
	"context"
	"errors"
	"fmt"
	"github.com/iamsayantan/konference"
	"gorm.io/gorm"
)

type roomRepo struct {
	db *gorm.DB
}

func (r *roomRepo) Store(ctx context.Context, room *konference.Room) error {
	if room.ID != 0 {
		return errors.New("invalid room details")
	}

	return r.db.WithContext(ctx).Create(&room).Error
}

func (r *roomRepo) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&konference.Room{}, id).Error
}

func (r *roomRepo) FindById(ctx context.Context, id uint) (*konference.Room, error) {
	return r.findByField(ctx, "id", id)
}

func (r *roomRepo) FindByInviteCode(ctx context.Context, inviteCode string) (*konference.Room, error) {
	return r.findByField(ctx, "invite_code", inviteCode)
}

func (r *roomRepo) findByField(ctx context.Context, fieldName string, fieldValue interface{}) (*konference.Room, error) {
	var room konference.Room
	err := r.db.WithContext(ctx).Where(
		fmt.Sprintf("%s = ?", fieldName),
		fieldValue,
	).First(&room).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &room, nil
}
func NewRoomRepository(db *gorm.DB) konference.RoomRepository {
	return &roomRepo{db: db}
}
