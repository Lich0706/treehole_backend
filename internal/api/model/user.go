package model

type CreateUserReq struct {
	//　User Entered Name
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
