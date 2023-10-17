package user

import (
	"github.com/gin-gonic/gin"
	"github.com/yuitaso/sampleWebServer/src/entity/user"
	userManager "github.com/yuitaso/sampleWebServer/src/entity/user/manager"
	"strconv"
)

type GetUserRequest struct {
	Id string `uri:"id" binding:"required"`
}

func GetOneById(c *gin.Context) {
	var request GetUserRequest
	var id int

	err := c.ShouldBindUri(&request)
	id, err = strconv.Atoi(request.Id)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()}) // いい感じに返すConfがあるはず
	}

	res := userManager.FindById(id)

	c.JSON(200, gin.H{"id": request.Id, "name": res.Name, "password": res.Password})
}

func Create(c *gin.Context) {
	err := userManager.Create(user.User{Name: "hander iikanji", Password: "hoge"})
	if err != nil {
		c.JSON(500, gin.H{"message": "できんかった"})
	}

	c.JSON(200, gin.H{"message": "できた"})
}
