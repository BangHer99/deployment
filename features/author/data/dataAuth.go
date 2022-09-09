package data

import (
	"alta/project2/features/author"

	"time"

	"gorm.io/gorm"
)

type User struct {
	Email     string `gorm:"primary key"`
	ID        uint   `gorm:"autoIncrement"`
	Name      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func toCore(user User) author.AuthCore {

	var core = author.AuthCore{
		ID:       int(user.ID),
		Email:    user.Email,
		Password: user.Password,
	}

	return core
}
