package models

import "time"

type Todo struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed" gorm:"default:false"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
