package model

type GetAuthReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
