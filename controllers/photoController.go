package controllers

import (
	"myGram/database"
	"myGram/helpers"
	"myGram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CreatePhoto godoc
// @Summary Post detail for a given id
// @Description Post Photo Detail corresponding to the input id
// @Tags photos
// @Accept json
// @Produce json
// @Param models.Photo body models.Photo true "create photo"
// @Success 200 {object} models.Photo
// @Router /photos/ [post]
func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	User := models.User{}
	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.User = &User

	if err := db.First(&User, "id = ?", userID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Gagal Mendapatkan Data"})
		return
	}

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Photo)
}

// GetPhotoById godoc
// @Summary Get detail for the given id
// @Description Get Photo Detail corresponding to the input id
// @Tags photos
// @Accept json
// @Produce json
// @Param id path int true "Id of the commment"
// @Success 200 {object} models.Photo
// @Router /photos/{id} [put]
func GetPhotoById(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}
	User := models.User{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.User = &User

	if err := db.First(&Photo, "id = ?", photoId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	userID := Photo.UserID

	err := db.First(&User, "id = ?", userID).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)
}

// UpdatePhoto godoc
// @Summary Update Photo identified by the given id
// @Description Update the photo corresponding to the input id
// @Tags photos
// @Accept json
// @Produce json
// @Param id path int true "Id of the commment to be updated"
// @Success 200 {object} models.Photo
// @Router /photos/{id} [put]
func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	err := db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoUrl: Photo.PhotoUrl}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)
}

// DeletePhoto godoc
// @Summary Delete Photo identified by the given id
// @Description Delete the photo corresponding to the input id
// @Tags photos
// @Accept json
// @Produce json
// @Param id path int true "Id of the commment to be deleted"
// @Success 200 {object} models.Photo
// @Router /photos/{id} [delete]
func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	Photo.ID = uint(photoId)
	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err := db.Where("id = ?", photoId).Delete(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, "Deleted")
}

// GetAllPhoto godoc
// @Summary Get Details
// @Description Get details of All Photo
// @Tags photos
// @Accept json
// @Produce json
// @Success 200 {object} models.Photo
// @Router /photos/ [get]
func GetAllPhoto(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	var Photo []models.Photo

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err := db.Find(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)
}
