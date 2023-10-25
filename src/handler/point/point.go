package point

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuitaso/sampleWebServer/src/handler"
	pointlogManager "github.com/yuitaso/sampleWebServer/src/manager/pointLog"
)

func FetchMyBalans(c *gin.Context) {

	authUser, ok := handler.GetAuthUserOrErrorRsponse(c)
	if !ok {
		return
	}

	balance, err := pointlogManager.GetSum(authUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot find.", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": balance})
}
