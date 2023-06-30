package handler

import (
	"TreeHole/treehole_backend/internal/api/model"
	"TreeHole/treehole_backend/internal/app"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListPosts(c *gin.Context) {
	// err = c.ShouldBind(&req)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, model.ErrorRes{Error: err.Error()})
	// }

	ret, err := app.PostDao.ListAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorRes{Error: err.Error()})
		return
	}

	response := model.GetRecordsResponse{}
	for _, item := range ret {
		var messageInterface map[string]interface{}
		messageByte, _ := json.Marshal(item)
		json.Unmarshal(messageByte, &messageInterface)
		response.Data = append(response.Data, messageInterface)
	}

	c.JSON(200, response)
}
