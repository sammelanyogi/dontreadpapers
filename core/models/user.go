package models

import (
	"log"

	database "github.com/sammelanyogi/dontreadpapers/services/database"
)

type User struct {
	ID        string `json:"user_id,omitempty"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	PhotoURL  string `json:"photo_url"`
	Active    bool   `json:"active,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

func (h User) GetByID(id string) (*User, error) {
	db := database.GetDB()
	user := new(User)
	err := db.Model(user).Where("id = ?", 1).Select()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return user, nil
}
