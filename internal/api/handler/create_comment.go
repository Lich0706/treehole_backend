package handler

import (
	"TreeHole/treehole_backend/internal/api/model"
	"TreeHole/treehole_backend/internal/app"
	dbModel "TreeHole/treehole_backend/internal/dao/model"
	"TreeHole/treehole_backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateComment(c *gin.Context) {
	var (
		err error
		req model.CreateCommentReq
	)
	err = c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorRes{Error: err.Error()})
		return
	}
	// Get token from header
	token := c.Request.Header.Get("token")
	claims, err := utils.ParseToken(token)

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorRes{Error: err.Error()})
		return
	}
	// Get user by encrypt info
	user, err := app.UserDao.FindByEncryptInfo(c, claims.EncrptedInfo)

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorRes{Error: err.Error()})
		return
	}

	newComment := dbModel.Comment{
		ReplyTo: req.ReplyToCid,
		PostID:  req.Pid,
		UserID:  user.ID,
		Content: req.Content,
	}

	// Check cuid and puid
	post, err := app.PostDao.GetPostById(c, req.Pid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorRes{Error: err.Error()})
		return
	}

	if post.UserID == newComment.UserID {
		newComment.NickName = utils.DzName
	} else {
		previousComment, err := app.CommentDao.GetCommentsByPidUid(c, req.Pid, user.ID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				count, err := app.CommentDao.CountExceptCommenter(c, req.Pid, post.UserID)
				if err != nil {
					c.JSON(http.StatusInternalServerError, model.ErrorRes{Error: err.Error()})
					return
				}
				newComment.NickName = utils.GenNickname(int(count))
			} else {
				c.JSON(http.StatusInternalServerError, model.ErrorRes{Error: err.Error()})
				return
			}
		} else {
			newComment.NickName = previousComment.NickName
		}
	}

	cid, err := app.CommentDao.CreateNewComment(c, newComment)

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorRes{Error: err.Error()})
		return
	}

	_, err = app.PostDao.UpdateReplyCountById(c, req.Pid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorRes{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.CreateObjResponse{
		ID: cid,
	})

}
