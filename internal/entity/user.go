package entity

type User struct {
	Base
	ProfilePicture *string `gorm:"type:varchar(255);default:null"`
	Name           string  `gorm:"type:varchar(100)"`
	Username       string  `gorm:"type:varchar(100);unique"`
	Email          string  `gorm:"type:varchar(255);unique"`
	Password       string  `gorm:"type:varchar(255)"`
}
