package api

type BasketItem struct {
	ID       uint     `json:"id" binding:"required"`
	Quantity int      `json:"quantity" binding:"required"`
	Product  *Product `json:"product,omitempty"`
	User     *User    `json:"user,omitempty"`
}
