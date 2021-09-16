package konference

import (
	"context"
	"errors"
	"github.com/bwmarrin/snowflake"
	"sync"
)

// MaxAllowedMembers is the maximum number that we allow per room.
const MaxAllowedMembers = 4
const invitationCodeBlockLength = 9999

var (
	idGenerator *snowflake.Node

	// ErrAlreadyInRoom is returned when an already existing member of the room tries to join again.
	ErrAlreadyInRoom = errors.New("user already in room")

	// ErrUserNotInRoom is returned when an non existing member is being removed from a room.
	ErrUserNotInRoom = errors.New("user does not belong to the room")

	//ErrRoomCapacityFull is triggered when someone tries to join after the room reaches its capacity.
	ErrRoomCapacityFull = errors.New("room capacity full")
)

func init() {
	var err error
	idGenerator, err = snowflake.NewNode(1)

	if err != nil {
		panic(err.Error())
	}
}

// generateInviteCode generates a random invite code for rooms.
func generateInviteCode(blockLength int) string {
	return idGenerator.Generate().String()
	//rand.Seed(time.Now().UnixNano())
	// format 9999-3343-3439
	//return strconv.Itoa(rand.Intn(blockLength)) + "-" + strconv.Itoa(rand.Intn(blockLength)) + "-" + strconv.Itoa(rand.Intn(blockLength))
}

// Room is the place where a call takes place.
type Room struct {
	ID         uint   `json:"id"`
	InviteCode string `json:"invite_code"`
	OwnerID    uint   `json:"-"`
	CreatedBy  *User  `json:"created_by" gorm:"foreignKey:OwnerID"`

	members map[uint]*User
	mu      sync.Mutex
}

// NewRoom creates a new Room. The invitation code is also generated here.
func NewRoom(owner *User) *Room {
	room := &Room{
		InviteCode: generateInviteCode(invitationCodeBlockLength),
		OwnerID:    owner.ID,
		CreatedBy:  owner,

		members: make(map[uint]*User, 0),
	}

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

func (r *Room) GetMember(memberId uint) *User {
	member, ok := r.members[memberId]
	if ok {
		return nil
	}

	return member
}

// AddMember adds an user to a room
func (r *Room) AddMember(member *User) error {
	r.mu.Lock()
	if r.IsFull() {
		return ErrRoomCapacityFull
	}

	if _, ok := r.members[member.ID]; ok {
		return ErrAlreadyInRoom
	}

	r.members[member.ID] = member
	r.mu.Unlock()

	return nil
}

// RemoveMember removes a member from the room.
// If the provided user is not found, then it is silently ignored.
func (r *Room) RemoveMember(memberId uint) error {
	r.mu.Lock()
	if _, ok := r.members[memberId]; !ok {
		return ErrUserNotInRoom
	}

	delete(r.members, memberId)
	r.mu.Unlock()

	return nil
}

// RoomRepository defines the methods to interact with the room storage.
type RoomRepository interface {
	Store(ctx context.Context, room *Room) error
	Delete(ctx context.Context, id uint) error
	FindById(ctx context.Context, id uint) (*Room, error)
	FindByInviteCode(ctx context.Context, inviteCode string) (*Room, error)
}

// RoomService defines the application methods exposed by the rooms domain.
type RoomService interface {
	// Create creates a new room. Creating a room does not add the creator as
	// a member by default.
	Create(ctx context.Context, ownerId uint) (*Room, error)

	// GetDetails returns room details by the invitation code.
	GetDetails(ctx context.Context, invitationCode string) (*Room, error)

	// Join adds a user to a room.
	Join(ctx context.Context, roomId uint, joiningUserId uint) error

	// Leave removes a user from the room.
	Leave(ctx context.Context, roomId uint, leavingUserId uint) error

	// IsResiding checks for if a particular user is a member of a room.
	IsResiding(ctx context.Context, roomId uint, userId uint) bool
}
