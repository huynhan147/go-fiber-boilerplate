package models

import "time"

type Transaction struct {
	ID        uint    `gorm:"primaryKey"`
	Amount    float64 `gorm:"not null"`
	Status    string  `gorm:"size:50"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
