package migration

import (
	Books "alta/project2/features/book/data"
	Users "alta/project2/features/user/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {

	db.AutoMigrate(&Users.User{})
	db.AutoMigrate(&Books.Book{})

}
