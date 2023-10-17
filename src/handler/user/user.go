package user

import (
	"github.com/gin-gonic/gin"
	"github.com/yuitaso/sampleWebServer/src/entity/user"
	userManager "github.com/yuitaso/sampleWebServer/src/manager/user"
	"net/http"
	"strconv"
)

type GetUserRequest struct { // TODO naming
	Id string `uri:"id" binding:"required"`
}

func GetOneById(c *gin.Context) {
	var request GetUserRequest // TODO naming requestはformdataで使ってるから別の名前にする
	var id int

	err := c.ShouldBindUri(&request)
	id, err = strconv.Atoi(request.Id)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()}) // いい感じに返すConfがあるはず
	}

	res, err := userManager.FindById(id)
	if err != nil {
		c.JSON(500, gin.H{"message": "cannot find.", "error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"id": request.Id, "name": res.Name, "password": res.Password})
}

type CreateRequest struct {
	Name     string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Create(c *gin.Context) {
	var request CreateRequest
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, err := userManager.Create(user.User{Name: request.Name, Password: request.Password})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"user_id": id})
}
