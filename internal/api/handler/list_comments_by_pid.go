package handler

import (
	"TreeHole/treehole_backend/internal/api/model"
	"TreeHole/treehole_backend/internal/app"
	dbModel "TreeHole/treehole_backend/internal/dao/model"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListCommentsByPid(c *gin.Context) {
	var (
		err error
		req model.GetCommentsReq
	)
	err = c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorRes{Error: err.Error()})
	}

	ret, err := app.CommentDao.GetCommentsByPid(c, req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorRes{Error: err.Error()})
		return
	}

	comments := nestComments(ret)

	response := model.GetRecordsResponse{}
	for _, item := range comments {
		if item.ReplyTo == 0 {
			var messageInterface map[string]interface{}
			messageByte, _ := json.Marshal(item)
			json.Unmarshal(messageByte, &messageInterface)
			response.Data = append(response.Data, messageInterface)
		}
	}

	c.JSON(200, response)
}

func nestComments(comments []*dbModel.Comment) []*dbModel.Comment {
	commentMap := make(map[int64]*dbModel.Comment)
	var roots []*dbModel.Comment

	// Initialize the map with all comments
	for i := range comments {
		commentMap[comments[i].ID] = comments[i]
	}

	// Build the tree
	for i := range comments {
		if comments[i].ReplyTo != 0 {
			parent, ok := commentMap[comments[i].ReplyTo]
			if ok {
				parent.Children = append(parent.Children, comments[i])
			}
		} else {
			roots = append(roots, comments[i])
		}
	}

	return roots
}
