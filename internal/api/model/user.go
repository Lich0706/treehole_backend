package model

type CreateUserReq struct {
	//　User Entered Name
	Name        string `json:"name"`
	Email       string `json:"email"`
	HashedEmail string `json:"hashed_email"`
	HashedPwd   string `json:"hashed_pwd"`
}

type LoginReq struct {
	Email     string `json:"email"`
	HashedPwd string `json:"hashed_pwd"`
}
