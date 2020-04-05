package users

// LoginRequest - a login json object to bind for access tokens
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
