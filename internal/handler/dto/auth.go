package dto

type Token string

type ErrorResponse struct {
	Message string `json:"message"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
type RegisterResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type DummyRequest struct {
	Role string `json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
