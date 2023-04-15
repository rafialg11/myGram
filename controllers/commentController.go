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

// CreateComment godoc
// @Summary Post detail for a given id
// @Description Post Comment Detail corresponding to the input id
// @Tags comments
// @Accept json
// @Produce json
// @Param models.Comment body models.Comment true "create comment"
// @Success 200 {object} models.Comment
// @Router /photos/{id}/comments [post]
func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	contentType := helpers.GetContentType(c)

	User := models.User{}
	Photo := models.Photo{}
	Comment := models.Comment{}

	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.PhotoID = uint(photoId)
	Comment.User = &User
	Comment.Photo = &Photo

	if err := db.First(&User, "id = ?", userID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Fail to get data"})
		return
	}

	if err := db.First(&Photo, "id = ?", photoId).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Fail to get data"})
		return
	}

	err := db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Comment)
}

// GetCommentById godoc
// @Summary Get detail for the given id
// @Description Get Comment Detail corresponding to the input id
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "Id of the commment"
// @Success 200 {object} models.Comment
// @Router /photos/{id}/comments/{id} [put]
func GetCommentById(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}
	User := models.User{}
	Photo := models.Photo{}
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	commentId, _ := strconv.Atoi(c.Param("commentId"))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.User = &User
	Comment.Photo = &Photo

	if err := db.First(&Comment, "id = ?", commentId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	userID := Comment.UserID

	if err := db.First(&Photo, "id = ?", photoId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	err := db.First(&User, "id = ?", userID).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)
}

// UpdateComment godoc
// @Summary Update Comment identified by the given id
// @Description Update the comment corresponding to the input id
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "Id of the commment to be updated"
// @Success 200 {object} models.Comment
// @Router /photos/{id}/comments/{id} [put]
func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.PhotoID = uint(photoId)
	Comment.ID = uint(commentId)

	err := db.Model(&Comment).Where("id = ?", commentId).Updates(models.Comment{Message: Comment.Message}).Error

	Comment.UserID = userID
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)
}

// DeleteComment godoc
// @Summary Delete Comment identified by the given id
// @Description Delete the comment corresponding to the input id
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "Id of the commment to be deleted"
// @Success 200 {object} models.Comment
// @Router /photos/{id}/comments/{id} [delete]
func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	Comment.ID = uint(commentId)
	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	err := db.Where("id = ?", commentId).Delete(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, "Deleted")
}

// GetAllComment godoc
// @Summary Get Details
// @Description Get details of All Comment
// @Tags comments
// @Accept json
// @Produce json
// @Success 200 {object} models.Comment
// @Router /photos/{id}/comments [get]
func GetAllComment(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	var Comment []models.Comment

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	err := db.Find(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)
}
