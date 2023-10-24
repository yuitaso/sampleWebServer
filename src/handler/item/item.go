package item

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yuitaso/sampleWebServer/src/entity"
	"github.com/yuitaso/sampleWebServer/src/handler"
	itemManager "github.com/yuitaso/sampleWebServer/src/manager/item"
)

type createRequest struct {
	Price int `form:"price" binding:"required"`
}

func Create(c *gin.Context) {
	authUser, ok := handler.GetAuthUserOrErrorRsponse(c)
	if !ok {
		return
	}

	fmt.Println(authUser)

	var request createRequest
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	fmt.Println("値段は？", request.Price)

	id, err := itemManager.Insert(authUser, request.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"uuid": id})
}

type editRequest struct {
	Price int `form:"price"`
}
type editUri struct {
	Uuid string `uri:"uuid" binding:"required"`
}

func Edit(c *gin.Context) {
	var uri editUri
	err := c.ShouldBindUri(&uri)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()}) // いい感じに返すConfがあるはず
		return
	}
	requested_uuid, err := uuid.Parse(uri.Uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Bad id format: %v", uri.Uuid)})
		return
	}

	var request editRequest
	err = c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	authUser, ok := handler.GetAuthUserOrErrorRsponse(c)
	if !ok {
		return
	}

	originalItem, err := itemManager.FindByUuid(requested_uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	fmt.Println("元データ", originalItem)
	fmt.Println("ユーザー", authUser)
	if authUser.Id != originalItem.UserId {
		c.JSON(http.StatusForbidden, gin.H{"message": "You do not have permission to edit this item."})
	}

	// ----
	var item entity.Item = *originalItem
	item.Price = request.Price

	err = itemManager.Update(&item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "updated.",
	})
}
