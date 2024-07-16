package libhttp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (pi *LibraryHttp) Ping(cnxt *gin.Context) {
	cpx := context.Background()
	err := pi.Libh.Ping(cpx)
	if err != nil {
		fmt.Println("3333", err)
		cnxt.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	cnxt.JSON(http.StatusOK, gin.H{
		"msg": "successful",
	})

}
