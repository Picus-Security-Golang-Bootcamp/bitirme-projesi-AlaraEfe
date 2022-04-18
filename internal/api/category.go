package api

type Category struct {
	ID           uint       `json:"category_id" binding:"required"`
	CategoryName string     `json:"category_name,omitempty"`
	Products     []*Product `json:"products"`
}

type CategorySlice []*Category
