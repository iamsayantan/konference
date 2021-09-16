package main

import (
	"fmt"
	"github.com/iamsayantan/konference/room"
	"github.com/iamsayantan/konference/server"
	"github.com/iamsayantan/konference/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"

	"github.com/iamsayantan/konference"
	mysqlSotrage "github.com/iamsayantan/konference/storage/mysql"
)

func main() {
	dsn := "root:12345@tcp(127.0.0.1:3306)/konference?charset=utf8mb4&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Could not connect to the database: %s", err.Error()))
	}

	err = db.AutoMigrate(&konference.User{})
	err = db.AutoMigrate(&konference.Room{})

	userRepo := mysqlSotrage.NewUserRepository(db)
	userService := user.NewUserService(userRepo)

	roomRepo := mysqlSotrage.NewRoomRepository(db)
	roomService := room.NewRoomService(roomRepo, userService)

	s := server.NewServer(userService, roomService)
	err = http.ListenAndServe(fmt.Sprintf(":%s", "8000"), s)
	if err != nil {
		log.Fatal(fmt.Sprintf("error starting the server: %s", err.Error()))
	}
}
