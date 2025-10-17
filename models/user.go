package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    Username   string     `gorm:"uniqueIndex" json:"username"`
    Password   string     `json:"password"`
    UserDetail UserDetail `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user_detail"`
}


// Enum สำหรับสถานะผู้ใช้
type UserStatus string

const (
    StatusActive   UserStatus = "active"
    StatusLeave    UserStatus = "leave"
    StatusInactive UserStatus = "inactive"
)

type UserDetail struct {
    ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    UserID uint `gorm:"uniqueIndex"` // ป้องกัน UserDetail ซ้ำ
    Code        string
    FirstName   string
    LastName    string
    Email       string
    Phone       string
    Position    string
    Department  string
    Status      UserStatus `gorm:"type:varchar(10)"`
}