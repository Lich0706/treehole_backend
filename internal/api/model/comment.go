package model

type CreateCommentReq struct {
	Pid        int64  `form:"pid,required" json:"pid"`
	ReplyToCid int64  `form:"reply_to_cid" json:"reply_to_cid"`
	Content    string `form:"content,required" json:"content"`
}

type GetCommentsReq struct {
	ID int64 `form:"id,required" json:"id"`
}
