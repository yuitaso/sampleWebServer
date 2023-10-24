package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuitaso/sampleWebServer/src/entity"
)

func GetAuthUserOrErrorRsponse(c *gin.Context) (*entity.User, bool) {
	val, exists := c.Get(entity.CtxAuthUserKey)
	if !exists {
		WhenInternalServerError(c, errors.New("Cannnot identify auth user."))
		return nil, exists
	}
	return val.(*entity.User), exists
}

func WhenInternalServerError(c *gin.Context, err error) {
	c.AbortWithStatusJSON(
		http.StatusInternalServerError,
		gin.H{"message": err.Error()},
	)
}
