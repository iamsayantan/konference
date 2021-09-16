package dto

import "github.com/iamsayantan/konference"

type RoomDetailsResponse struct {
	RoomDetails konference.Room `json:"room_details"`
}
