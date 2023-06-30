package model

type GetAuthReq struct {
	Email     string `json:"email"`
	HashedPwd string `json:"hashed_pwd"`
}
