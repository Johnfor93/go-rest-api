package contactcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/Johnfor93/go-rest-api/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"

	"fmt"
	"strconv"
)

func ShowAll(c *gin.Context) {

	query := c.Request.URL.Query()
	limit := 100
	offset:= 0

	var contacts []models.Contact

	fmt.Printf("%T %v", query.Get("offset"), query.Get("offset"))

	if query["offset"] != nil{
		offset, _ = strconv.Atoi(query.Get("offset"))
	}

	if query["limit"] != nil{
		limit, _ = strconv.Atoi(query.Get("limit"))
	}

	if query["filter"] != nil{
		fmt.Printf("masuk sini")
		models.DB.Where("gender = ?", query.Get("filter")).Limit(limit).Offset(offset).Find(&contacts)
	} else {
		models.DB.Limit(limit).Offset(offset).Find(&contacts)
	}

	c.JSON(http.StatusOK, gin.H{"contacts": contacts})

}

func Find(c *gin.Context) {
	var contact models.Contact
	id := c.Param("id")

	if err := models.DB.First(&contact, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"contact": contact})
}

func Insert(c *gin.Context) {

	var contact models.Contact

	if err := c.ShouldBindJSON(&contact); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&contact)

	// fmt.Printf("%v\n", contact)
	c.JSON(http.StatusOK, gin.H{"contact": contact})
}

func Update(c *gin.Context) {
	var contact models.Contact
	id := c.Param("id")

	if err := c.ShouldBindJSON(&contact); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&contact).Where("id = ?", id).Updates(&contact).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})

}

func Delete(c *gin.Context) {

	var contact models.Contact

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&contact, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus contact"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}