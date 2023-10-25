package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yuitaso/sampleWebServer/src/auth"
	"github.com/yuitaso/sampleWebServer/src/entity"
	userManager "github.com/yuitaso/sampleWebServer/src/manager/user"
)

func Create(c *gin.Context) {
	var request createRequest
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	newUser, err := userManager.CreateUserWithPointGrant(request.Email, request.Password, entity.InitialPointAmount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success.",
		"user": map[string]string{
			"uuid":  newUser.Uuid.String(),
			"email": newUser.Email,
		}})
}

type getOneByIdUri struct {
	Uuid string `uri:"uuid" binding:"required"`
}

func FetchOneByUuid(c *gin.Context) {
	var uri getOneByIdUri
	err := c.ShouldBindUri(&uri)

	user_uuid, err := uuid.Parse(uri.Uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id."}) // いい感じに返すConfがあるはず
	}

	result, err := userManager.FindByUuid(&user_uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot find.", "error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"id": result.Id, "uuid": result.Uuid, "email": result.Email})
}

func FetchMe(c *gin.Context) {
	var user *entity.User
	if val, exists := c.Get(entity.CtxAuthUserKey); !exists {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "ise."})
		return
	} else {
		user = val.(*entity.User)
	}

	c.JSON(http.StatusOK, gin.H{"uuid": user.Uuid, "email": user.Email})
}

type createRequest struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
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
