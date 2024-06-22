package model

type CreatePostReq struct {
	Content string
}

type GetPostReq struct {
	ID int64 `form:"id,required" json:"id"`
}
