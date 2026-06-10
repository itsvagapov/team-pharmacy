package models

type Review struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	UserID     uint   `json:"user_id"`
	MedicineID uint   `json:"medicine_id"`
	Rating     int    `json:"rating"`
	Text       string `json:"text"`

	User string

	Medicine Medicine `json:"-"`
}

type ReviewCreateRequest struct {
	UserID uint   `json:"user_id"`
	Rating int    `json:"rating"`
	Text   string `json:"text"`
}

type ReviewUpdateRequest struct {
	Text *string `json:"text"`
}
