package account

type RegisterRequest struct {
	Username string
	Password string
	Role     string
}
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Role     string `json:"role" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Response struct {
	Code    int    `json:"code" binding:"required"`
	Message string `json:"message" binding:"required"`
	Data    any    `json:"data" binding:"required"`
}

type StatusRequest struct {
	Status string `json:"status" binding:"required"`
}
