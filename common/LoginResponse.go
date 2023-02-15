package common

type LoginResponse struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}
