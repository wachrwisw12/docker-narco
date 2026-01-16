package models

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	RoleId   int8   `json:"role_id"`
	Status   string `json:"status"`
}
