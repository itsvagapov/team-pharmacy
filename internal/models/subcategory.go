package models

type Subcategory struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	CategoryID uint   `json:"category_id"`
	Name       string `json:"name"`
	Category Category
}

type SubcategoryCreateRequest struct {
	Name string `json:"name"`
}

type SubcategoryUpdateRequest struct {
	Name *string `json:"name"`
}
