package db

import(
	"github.com/gin-gonic/gin"
	"github.com/satyasiba/library/model"
)


func getBodyDate(c *gin.Context){
	var input Member

	body := c.Request.body
	values, _ := ioutils.ReadAll(body)
	c.JSON(200,gin.H{
		"body_data": string(value),
	})
}

func getQueryParam(c *gin.Context){
	id := c.Query("id")
	c.JSON(200,gin.H{
		"members_id": id,
	})
}

func getDatafromURL(c *gin.Context){
	id := c.Param("id")
	c.JSON(200, gin.H{
		"members_id": id,
	})
}

func insertIntoDB(){
	
}

