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
    Salary       float32
    TimeWork []TimeWork
    Leave []Leave
}

type TimeWork struct {
    ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    UserDetail uint
    TimeIn float32
    TimeOut float32
    TotalWorkTime string
    Date string `gorm:"type:DATE"`
}

type LeaveStatus string

const (
    StatusPending LeaveStatus = "pending"
    StatusApproved LeaveStatus = "approved"
    StatusRejected LeaveStatus = "rejected"
)
type Leave struct {
    ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    UserDetail uint
    StartDate string `gorm:"type:DATE"`
    EndDate string `gorm:"type:DATE"`
    Description string
    Total float32
    Status LeaveStatus `gorm:"type:varchar(10)"`
}