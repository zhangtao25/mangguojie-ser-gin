package api

import (
	"github.com/zhangtao25/mangostreet-ser-gin/service/note_service"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)


func GetNote(c *gin.Context) {
	title := c.Param("title")

	noteService := note_service.Note{Title: title}
	note, err := noteService.Get()


	if err != nil {
		c.JSON(200,gin.H{
			"data":"err",
		})
		return
	}
	c.JSON(200,gin.H{
		"data":note,
	})
	return
}