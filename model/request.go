package model

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type ShortingURLRequest struct {
	OriginalURL string `json:"original_url"`
}
