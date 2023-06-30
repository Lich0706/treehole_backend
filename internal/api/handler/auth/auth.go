package auth

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

func GetAuth(c *gin.Context) {
	var (
		err error
		req model.GetAuthReq
	)
	err = c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorRes{Error: err.Error()})
		return
	}

	email := req.Email
	hashedPwd := req.HashedPwd

	encryptedEmail, err := utils.AESEncrypted(email, hashedPwd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorRes{Error: err.Error()})
		return
	}

	user, err := app.UserDao.FindOne(c, dbModel.User{
		EncrptedEmail: encryptedEmail,
	})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(-1, model.ErrorRes{
				Error: errorsMsg.FAILED_AUTH,
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, model.ErrorRes{Error: err.Error()})
			return
		}
	}

	if user.Role == dbModel.BannedUser {
		c.JSON(-1, model.ErrorRes{
			Error: errorsMsg.BANNED_USER,
		})
		return
	}

	token, err := utils.GenerateToken(user.Name, user.EncrptedEmail)
	if err != nil {
		c.JSON(-1, model.ErrorRes{
			Error: errorsMsg.ERROR_GEN_TOKEN,
		})
		return
	}

	c.JSON(200, model.GetAuthResponse{
		Token: token,
	})
}
