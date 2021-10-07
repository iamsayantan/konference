package handlers

import (
	"github.com/go-chi/chi"
	"github.com/iamsayantan/konference"
	"github.com/iamsayantan/konference/server/dto"
	"github.com/iamsayantan/konference/server/middlewares"
	"github.com/iamsayantan/konference/server/rendering"
	"net/http"
	"time"
)

type roomHandler struct {
	service konference.RoomService
}

func (rh *roomHandler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(middlewares.AuthChecker)

	r.Post("/", rh.createRoom)
	return r
}

func (rh *roomHandler) createRoom(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second)
	loggedInUserID := middlewares.GetAuthenticatedUserIdFromContext(r.Context())
	room, err := rh.service.Create(r.Context(), loggedInUserID)
	if err != nil {
		rendering.RenderError(w, r, err.Error(), "room.createRoom.room_creation_error", http.StatusBadRequest)
		return
	}

	resp := dto.RoomDetailsResponse{RoomDetails: *room}
	rendering.RenderSuccessWithData(w, r, "success", http.StatusCreated, resp)
}

//func (rh *roomHandler) findByCode(w http.ResponseWriter, r *http.Request) {
//	roomCode := chi.URLParam(r, "inviteCode")
//	room := rh.service.GetDetails()
//}

func NewRoomHandler(s konference.RoomService) Handler {
	return &roomHandler{service: s}
}
