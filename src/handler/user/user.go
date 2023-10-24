package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yuitaso/sampleWebServer/src/auth"
	"github.com/yuitaso/sampleWebServer/src/entity"
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
	c.JSON(http.StatusOK, gin.H{"id": uri.Id, "id_hash": res.Uuid, "email": res.Email})
}

func GetUserMe(c *gin.Context) {
	var user *entity.User
	if val, exists := c.Get(entity.CtxAuthUserKey); !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "ise."})
		return
	} else {
		user = val.(*entity.User)
	}

	c.JSON(http.StatusOK, gin.H{"uuid": user.Uuid, "email": user.Email})
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

	id, err := userManager.Insert(request.Email, request.Password)
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

	user, err := userManager.VerifyAndGetUser(request.Email, request.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized."})
		return
	}

	token, err := auth.GenerateToken(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Faild to issue token.")})
		return
	}

	c.Header("X-Token", token)
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
