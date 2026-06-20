package models

import "time"

type Review struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	UserID     uint      `json:"user_id"`
	MedicineID uint      `json:"medicine_id"`
	Rating     int       `json:"rating"`
	Text       string    `json:"text"`
	CreatedAt  time.Time `json:"created_at"`

	Medicine Medicine `json:"-"`
}

type ReviewCreateRequest struct {
	UserID uint   `json:"user_id"`
	Rating int    `json:"rating"`
	Text   string `json:"text"`
}

type ReviewUpdateRequest struct {
	Rating *int    `json:"rating"`
	Text   *string `json:"text"`
}
