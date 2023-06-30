package handler

import (
	"TreeHole/treehole_backend/internal/api/errorsMsg"
	"TreeHole/treehole_backend/internal/api/model"
	"TreeHole/treehole_backend/internal/app"
	dbModel "TreeHole/treehole_backend/internal/dao/model"
	"TreeHole/treehole_backend/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUser(c *gin.Context) {
	var (
		err error
		req model.CreateUserReq
	)
	err = c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorRes{Error: err.Error()})
		return
	}

	email := req.Email
	hashedPwd := req.HashedPwd
	hashedEmail := req.HashedEmail

	encryptedEmail, err := utils.AESEncrypted(email, hashedPwd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorRes{Error: err.Error()})
		return
	}

	err = app.HashedEmailDao.FindOne(c, hashedEmail)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = app.HashedEmailDao.Create(c, hashedEmail)
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.ErrorRes{Error: err.Error()})
			return
		}
	} else {
		// 该邮箱已经被注册了
		c.JSON(-1, model.ErrorRes{Error: errorsMsg.ERROR_EXIST_EMAIL})
		return
	}

	newUser := dbModel.User{
		Name:          req.Name,
		Role:          dbModel.NormalUser,
		EncrptedEmail: encryptedEmail,
	}
	err = app.UserDao.CreateNormalUser(c, newUser)

	if err != nil {
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.ErrorRes{Error: err.Error()})
			return
		}
	}

	c.JSON(200, gin.H{
		"message": "successfully created user",
	})
}
