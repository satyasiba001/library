package libhttp

import (
	"github.com/gin-gonic/gin"
	"github.com/satyasiba001/library/services"
)

type LibraryHttp struct {
	Engine *gin.Engine
	Libh   *services.LibraryHandler
}

func NewLibraryHttp(eng *gin.Engine,lh *services.LibraryHandler)*LibraryHttp{
	return &LibraryHttp{
		Engine: eng,
		Libh: lh,
	}
}


func (lh *LibraryHttp) Routes() {
	routes := lh.Engine.Group("/library")
	routes.GET("/ping", lh.Ping)
}

