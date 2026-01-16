package models

type AuthRequest struct {
	Uid      string `json:"uid"`
	Username string `json:"username"`
	Password string `json:"password"`
}
