package konference

import (
	"context"
	"errors"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// MaxAllowedMembers is the maximum number that we allow per room.
const MaxAllowedMembers = 4
const invitationCodeBlockLength = 9999

var (
	// ErrAlreadyInRoom is returned when an already existing member of the room tries to join again.
	ErrAlreadyInRoom = errors.New("user already in room")

	//ErrRoomCapacityFull is triggered when someone tries to join after the room reaches its capacity.
	ErrRoomCapacityFull = errors.New("room capacity full")
)

// generateInviteCode generates a random invite code for rooms.
func generateInviteCode(blockLength int) string {
	rand.Seed(time.Now().UnixNano())
	// expected format 9999-3343-3439
	return strconv.Itoa(rand.Intn(blockLength)) + "-" + strconv.Itoa(rand.Intn(blockLength)) + "-" + strconv.Itoa(rand.Intn(blockLength))
}

// Room is the place where a call takes place.
type Room struct {
	ID         uint   `json:"id"`
	InviteCode string `json:"inviteCode"`
	CreatedBy  *User  `json:"createdBy"`

	members map[uint]*User
	mu      sync.Mutex
}

// WaitingRoom is the place where users are added when they want to join a room.
// After someone approves the user they would be moved as the room member.
type WaitingRoom struct {
	ID   uint  `json:"id"`
	Room *Room `json:"room"`

	waitingMembers map[uint]*User
	mu             sync.Mutex
}

// NewRoom creates a new Room. The owner will be automatically added after the room
// is created. The invitation code is also generated here.
func NewRoom(owner *User) *Room {
	room := &Room{
		InviteCode: generateInviteCode(invitationCodeBlockLength),
		CreatedBy:  owner,

		members: make(map[uint]*User, 0),
	}

	_ = room.AddMember(owner)
	return room
}

// Members returns the all the members in the room.
func (r *Room) Members() []*User {
	users := make([]*User, 0)
	for _, user := range r.members {
		users = append(users, user)
	}
	return users
}

// MemberCount returns the total number of members currently present in the room.
func (r *Room) MemberCount() int {
	return len(r.members)
}

// IsFull returns if the room capacity reached the maximum allowed limit.
func (r *Room) IsFull() bool {
	return r.MemberCount() >= MaxAllowedMembers
}

// RefreshInviteCode updates the current invitation code.
func (r *Room) RefreshInviteCode() {
	r.InviteCode = generateInviteCode(invitationCodeBlockLength)
}

// AddMember adds an user to a room
func (r *Room) AddMember(user *User) error {
	r.mu.Lock()
	if r.IsFull() {
		return ErrRoomCapacityFull
	}

	if _, ok := r.members[user.ID]; ok {
		return ErrAlreadyInRoom
	}

	r.members[user.ID] = user
	r.mu.Unlock()

	return nil
}

// RoomRepository defines the methods to interact with the room storage.
type RoomRepository interface {
	Store(ctx context.Context, room *Room) (*Room, error)
	Update(ctx context.Context, room *Room) (*Room, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (*Room, error)
	FindByInviteCode(ctx context.Context, inviteCode string) (*User, error)
}
