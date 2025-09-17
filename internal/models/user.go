package models

import (
	"food-delivery-api/constants/user_role"
	"time"
)

type User struct {
	UserID      uint               `gorm:"primaryKey;autoIncrement;column:userID" json:"userID"`
	Name        string             `gorm:"type:varchar(100);not null;column:name" json:"name"`
	Email       string             `gorm:"type:varchar(100);unique;not null;column:email" json:"email"`
	Password    string             `gorm:"type:varchar(255);not null;column:password" json:"-"`
	UserRole    user_role.UserRole `gorm:"type:enum('customer','restaurant','driver','admin');default:'customer';not null;column:userRole" json:"userRole"`
	CreatedDate time.Time          `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:createdDate" json:"createdDate"`
	UpdatedDate time.Time          `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP;column:updatedDate" json:"updatedDate"`
}

func (User) TableName() string {
	return "ms_user"
}
