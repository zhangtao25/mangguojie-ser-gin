package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"test/models"
)


func errorHandler(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func GetNotes(c *gin.Context) {
	res := models.GetNotes()
	c.JSON(200, gin.H{
		"notes": res,
	})
	return
}

func GetNote(c *gin.Context) {
	id := c.Params.ByName("id")
	res := models.GetNote(id)
	c.JSON(200, gin.H{
		"notes": res,
	})
	return
}
