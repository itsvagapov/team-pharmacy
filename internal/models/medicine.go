package models

type Medicine struct {
	ID                   uint    `gorm:"primaryKey" json:"id"`
	Name                 string  `json:"name"`
	Description          string  `json:"description"`
	Price                float64 `json:"price"`
	InStock              bool    `json:"in_stock"`
	StockQuantity        int     `json:"stock_quantity"`
	CategoryID           uint    `json:"category_id"`
	SubcategoryID        uint    `json:"subcategory_id"`
	Manufacturer         string  `json:"manufacturer"`
	PrescriptionRequired bool    `json:"prescription_required"`
	AvgRating            float64 `json:"avg_rating"`

	Category    Category    `json:"-"`
	Subcategory Subcategory `json:"-"`
}

type MedicineCreateRequest struct {
	Name                 string  `json:"name"`
	Description          string  `json:"description"`
	Price                float64 `json:"price"`
	InStock              bool    `json:"in_stock"`
	StockQuantity        int     `json:"stock_quantity"`
	CategoryID           uint    `json:"category_id"`
	SubcategoryID        uint    `json:"subcategory_id"`
	Manufacturer         string  `json:"manufacturer"`
	PrescriptionRequired bool    `json:"prescription_required"`
}

type MedicineUpdateRequest struct {
	Price         *float64 `json:"price"`
	InStock       *bool    `json:"in_stock"`
	StockQuantity *int     `json:"stock_quantity"`
}

type MedicineFilter struct {
	Search        string `form:"search"`
	CategoryID    *uint  `form:"category_id"`
	SubcategoryID *uint  `form:"subcategory_id"`
	InStock       *bool  `form:"in_stock"`
}
