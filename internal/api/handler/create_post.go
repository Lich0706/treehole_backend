package handler

import (
	"TreeHole/treehole_backend/internal/api/model"
	"TreeHole/treehole_backend/internal/app"
	dbModel "TreeHole/treehole_backend/internal/dao/model"
	"TreeHole/treehole_backend/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var (
		err error
		req model.CreatePostReq
	)
	err = c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorRes{Error: err.Error()})
		return
	}

	token := c.Request.Header.Get("token")
	claims, err := utils.ParseToken(token)

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorRes{Error: err.Error()})
		return
	}

	user, err := app.UserDao.FindByEncryptInfo(c, claims.EncrptedInfo)

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorRes{Error: err.Error()})
		return
	}

	newPost := dbModel.Post{
		UserID:    user.ID,
		Content:   req.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	postId, err := app.PostDao.CreateNewPost(c, newPost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorRes{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.CreateObjResponse{
		ID: postId,
	})

}
