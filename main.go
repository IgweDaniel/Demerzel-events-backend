package main

import (
	"demerzel-events/api"
	"demerzel-events/configs"
	"demerzel-events/internal/db"
	"demerzel-events/internal/models"
	"fmt"
	"os"
	"strconv"
)

func Seed() {
	users := []models.User{
		{Email: "test@test.com", Name: "Joe420"},
		{Email: "yes@yes.com", Name: "Bob"},
	}
	db.DB.Create(&users)

	groups := []models.Group{
		{Name: "laughter and sorrows group"},
		{Name: "dancing group"},
	}

	db.DB.Create(&groups)

	usergroups := []models.UserGroup{
		{UserID: users[0].Id, GroupID: groups[0].ID},
		{UserID: users[0].Id, GroupID: groups[1].ID},
		{UserID: users[1].Id, GroupID: groups[0].ID},
		{UserID: users[1].Id, GroupID: groups[1].ID},
	}

	db.DB.Create(&usergroups)
}

func main() {
	configs.Load()
	Seed()
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		panic(fmt.Sprintf("Failed to conver PORT to integer: %v", err))
	}

	srv := api.NewServer(uint16(port), api.BuildRoutesHandler())
	srv.Listen()
}
