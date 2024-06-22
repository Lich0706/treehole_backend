package handler

import (
	"TreeHole/treehole_backend/internal/api/model"
	"TreeHole/treehole_backend/internal/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPost(c *gin.Context) {
	var (
		err error
		req model.GetPostReq
	)
	err = c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorRes{Error: err.Error()})
	}

	ret, err := app.PostDao.GetPostById(c, req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorRes{Error: err.Error()})
		return
	}

	c.JSON(200, ret)
}
