package api

type Product struct {
	ID               uint      `json:"product_id" binding:"required"`
	ProductName      string    `json:"product_name,omitempty"`
	ProductStock     int       `json:"product_stock,omitempty"`
	ProductPrice     float64   `json:"product_price,omitempty"`
	ProductStockCode string    `json:"product_stock_code,omitempty"`
	ProductType      string    `json:"product_type,omitempty"`
	ProductColor     string    `json:"product_color,omitempty"`
	ProductSize      string    `json:"product_size,omitempty"`
	ProductMaterial  string    `json:"product_material,omitempty"`
	Gender           string    `json:"gender,omitempty"`
	Category         *Category `json:"category,omitempty"`
}
