package model

type User struct {
	Id        int64  `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Username  string `json:"username"`
	Password  string `json:"-"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Image     string `json:"image"`
}
