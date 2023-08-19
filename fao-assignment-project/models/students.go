package models

import "time"



type Students struct {
	ID        uint     `gorm:"primaryKey"`
	Name      string   `gorm:"not null;type:varchar(191)"`
	Age       int      `gorm:"not null"`
	Scores    []Scores `gorm:"foreignkey:StudentID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
