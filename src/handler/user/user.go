package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	userManager "github.com/yuitaso/sampleWebServer/src/manager/user"
)

type GetOneByIdUri struct {
	Id string `uri:"id" binding:"required"`
}

func GetOneById(c *gin.Context) {
	var uri GetOneByIdUri
	var id int

	err := c.ShouldBindUri(&uri)
	id, err = strconv.Atoi(uri.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()}) // いい感じに返すConfがあるはず
	}

	res, err := userManager.FindById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot find.", "error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"id": uri.Id, "id_hash": res.IdHash, "email": res.Email})
}

type CreateRequest struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Create(c *gin.Context) {
	var request CreateRequest
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, err := userManager.Create(request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_id": id})
}

type AuthRequest struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Authenticate(c *gin.Context) {
	var request AuthRequest
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// if err := userManager.Authenticate(request.Email, request.Password); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	// 	return
	// }

	c.Header("X-Token", "hogehoge")
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
