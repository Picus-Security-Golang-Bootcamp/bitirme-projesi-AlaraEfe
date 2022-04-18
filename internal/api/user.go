package api

type User struct {
	ID       uint          `json:"user_id" binding:"required"`
	Email    string        `json:"email,omitempty"`
	Password string        `json:"password,omitempty"`
	Role     string        `json:"role,omitempty"`
	Items    []*BasketItem `json:"items,omitempty"`
}
