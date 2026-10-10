package models

import "time"

type User struct {
    ID           uint      `json:"id" gorm:"primaryKey"`
    Name         string    `json:"name" gorm:"not null"`
    Email        string    `json:"email" gorm:"uniqueIndex;not null"`
    PasswordHash string    `json:"-" gorm:"not null"`
    CreatedAt    time.Time `json:"created_at"`
}