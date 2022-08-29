package model

import "time"

type Todo struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Title     string    `json:"title" validate:"required"`
	Number    int       `json:"number" gorm:"autoIncrement:true"`
	Timestamp time.Time `json:"timestamp"`
	Content   string    `json:"content" validate:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type TodoHTTPQueryParameter struct {
	ID        string `json:"id,omitempty" query:"id"`
	Timestamp string `json:"timestamp,omitempty" query:"timestamp"`
	Title     string `json:"title,omitempty" query:"title"`
	Number    int    `json:"number,omitempty" query:"number"`
}
