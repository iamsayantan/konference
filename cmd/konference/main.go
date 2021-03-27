package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/iamsayantan/konference"
)

func main() {
	dsn := "root:12345@tcp(127.0.0.1:3306)/konference?charset=utf8mb4"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Could not connect to the database: %s", err.Error()))
	}

	err = db.AutoMigrate(&konference.User{})

	user := konference.NewUser(
		"sayantan.das@codelogicx.com",
		"Sayantan",
		"Das",
		"sayantan94",
	)

	room := konference.NewRoom(user)
	fmt.Printf("created invite code: %s, current members: %d\n", room.InviteCode, room.MemberCount())
	room.RefreshInviteCode()
	fmt.Printf("updated invite code: %s\n", room.InviteCode)
}
