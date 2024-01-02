package controllers

import (
	"net/http"

	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"rest-api/app/PhotoResponse"
	"rest-api/helpers/formatter"
	"rest-api/models"
)

type PhotoController struct {
	db *gorm.DB
}

func NewPhotoController(db *gorm.DB) *PhotoController {
	return &PhotoController{db}
}

func (h *PhotoController) Get(c *gin.Context) {
	var userPhoto models.Photo
	if err := h.db.Preload("User").Find(&userPhoto).Error; err != nil {
		handlePhotoError(c, http.StatusBadRequest, "Failed to Get Your Photo", err)
		return
	}

	if userPhoto.PhotoURL == "" {
		handlePhotoError(c, http.StatusBadRequest, "Please Upload Your Photo", nil)
		return
	}

	formatterr := photo.FormatPhoto(&userPhoto, "")
	response := formatter.ApiResponse(http.StatusOK, "success", formatterr, "Successfully Fetch User Photo")
	c.JSON(http.StatusOK, response)
}

func (h *PhotoController) Create(c *gin.Context) {
	var userPhoto models.Photo
	var countPhoto int64
	currentUser := c.MustGet("currentUser").(models.User)

	h.db.Model(&userPhoto).Where("user_id = ?", currentUser.ID).Count(&countPhoto)
	if countPhoto > 0 {
		handlePhotoError(c, http.StatusBadRequest, "You Already Have a Photo", nil)
		return
	}

	var input models.Photo
	if err := c.ShouldBind(&input); err != nil {
		handleValidationError(c, http.StatusUnprocessableEntity, err, "Failed to Upload User Photo")
		return
	}

	file, err := c.FormFile("photo_profile")
	if err != nil {
		handleValidationError(c, http.StatusUnprocessableEntity, err, "Failed to Upload User Photo")
		return
	}

	extension := file.Filename
	path := "static/images/" + uuid.New().String() + extension

	if err := c.SaveUploadedFile(file, path); err != nil {
		handlePhotoError(c, http.StatusUnprocessableEntity, "Failed to Upload User Photo", err)
		return
	}

	h.InsertPhoto(input, path, currentUser.ID)

	data := gin.H{"is_uploaded": true}
	response := formatter.ApiResponse(http.StatusOK, "success", data, "Photo Profile Successfully Uploaded")
	c.JSON(http.StatusOK, response)
}

func (h *PhotoController) InsertPhoto(userPhoto models.Photo, fileLocation string, currUserID int) error {
	savePhoto := models.Photo{
		UserID:   currUserID,
		Title:    userPhoto.Title,
		Caption:  userPhoto.Caption,
		PhotoURL: fileLocation,
	}

	if err := h.db.Debug().Create(&savePhoto).Error; err != nil {
		return err
	}
	return nil
}

func (h *PhotoController) Update(c *gin.Context) {
	var userPhoto models.Photo
	currentUser := c.MustGet("currentUser").(models.User)

	if err := h.db.Where("user_id = ?", currentUser.ID).Find(&userPhoto).Error; err != nil {
		handlePhotoError(c, http.StatusBadRequest, "Photo Profile Failed to Update", err)
		return
	}

	var input models.Photo
	if err := c.ShouldBind(&input); err != nil {
		handleValidationError(c, http.StatusBadRequest, err, "Photo Profile Failed to Update")
		return
	}

	file, err := c.FormFile("update_profile")
	if err != nil {
		handleValidationError(c, http.StatusUnprocessableEntity, err, "Failed to Update User Photo")
		return
	}

	extension := file.Filename
	path := "static/images/" + uuid.New().String() + extension

	if err := c.SaveUploadedFile(file, path); err != nil {
		handlePhotoError(c, http.StatusBadRequest, "Photo Profile Failed to Upload", err)
		return
	}

	h.UpdatePhoto(input, &userPhoto, path)

	data := photo.FormatPhoto(&userPhoto, "regular")
	response := formatter.ApiResponse(http.StatusOK, "success", data, "Photo Profile Successfully Updated")
	c.JSON(http.StatusOK, response)
}

func (h *PhotoController) UpdatePhoto(oldPhoto models.Photo, newPhoto *models.Photo, path string) error {
	newPhoto.Title = oldPhoto.Title
	newPhoto.Caption = oldPhoto.Caption
	newPhoto.PhotoURL = path

	if err := h.db.Save(newPhoto).Error; err != nil {
		return err
	}
	return nil
}

func (h *PhotoController) Delete(c *gin.Context) {
	var userPhoto models.Photo
	currentUser := c.MustGet("currentUser").(models.User)

	if err := h.db.Where("user_id = ?", currentUser.ID).Delete(&userPhoto).Error; err != nil {
		handlePhotoError(c, http.StatusBadRequest, "Failed to delete user photo", err)
		return
	}

	data := gin.H{"is_deleted": true}
	response := formatter.ApiResponse(http.StatusOK, "success", data, "User Photo Successfully Deleted")
	c.JSON(http.StatusOK, response)
}

// Helper function to handle photo errors
func handlePhotoError(c *gin.Context, statusCode int, message string, err error) {
	response := formatter.ApiResponse(statusCode, "error", nil, message)
	c.JSON(statusCode, response)
}

// Helper function to handle validation errors
func handleValidationError(c *gin.Context, statusCode int, err error, message string) {
	errors := formatter.FormatValidationError(err)
	errorMessages := gin.H{"errors": errors}
	response := formatter.ApiResponse(statusCode, "error", errorMessages, message)
	c.JSON(statusCode, response)
}
