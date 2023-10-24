package item

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuitaso/sampleWebServer/src/entity"
	itemManager "github.com/yuitaso/sampleWebServer/src/manager/item"
)

type CreateRequest struct {
	Price int `form:"price" binding:"required"`
}

func Create(c *gin.Context) {
	var user *entity.User
	if val, exists := c.Get(entity.CtxAuthUserKey); !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Cannot identigy request user."})
		return
	} else {
		user = val.(*entity.User)
	}

	fmt.Println(user)

	var request CreateRequest
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	fmt.Println("値段は？", request.Price)

	id, err := itemManager.Insert(user, request.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"uuid": id})
}
