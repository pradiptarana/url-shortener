package model

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type URL struct {
	Id          int    `db:"id"`
	OriginalURL string `db:"original_url"`
	ShortURL    string `db:"short_url"`
	CreatedAt   int64  `db:"created_at"`
}
