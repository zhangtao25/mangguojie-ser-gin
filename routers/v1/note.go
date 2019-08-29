package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io"
	"log"
	"net/http"
	"os"
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
func AddNote(c *gin.Context)  {

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}

	filename := header.Filename
	out, err := os.Create("public/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}

	filepath := "http://localhost:8080/file/" + filename
	c.JSON(http.StatusOK, gin.H{"filepath": filepath})
	return
}
