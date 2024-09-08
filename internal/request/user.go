package request

type RegisterUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LoginRequest struct {
	Email string `json:"email"`
}
