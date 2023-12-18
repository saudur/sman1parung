package models

import "time"

type User struct {
	Id        int64     `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	Createdat time.Time `json:"created_at" gorm:"type:timestamp;not null;default:now()"`
	Updatedat time.Time `json:"updated_at" gorm:"type:timestamp;not null;default:now()"`
}
