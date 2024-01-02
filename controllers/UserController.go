package controllers

import (
	"net/http"
	"encoding/json"

	"github.com/gin-gonic/gin"
	UserResponse "rest-api/app/UserResponse"
	"rest-api/helpers/formatter"
	"rest-api/helpers/bcrypt"
	"rest-api/helpers/jwt"
	UserModel "rest-api/models"
	"gorm.io/gorm"
)

type userController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *userController {
	return &userController{db}
}

func (h *userController) Register(c *gin.Context) {
	var user UserModel.User

	if err := c.ShouldBindJSON(&user); err != nil {
		response := formatter.ApiResponse(http.StatusBadRequest, "error", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	hashedPassword, err := bcrypt.HashPassword(user.Password)
	if err != nil {
		// Handle the error and return an error response.
		handleRegistrationError(c, err)
		return
	}

	user.Password = hashedPassword

	if err := h.db.Debug().Create(&user).Error; err != nil {
		// Handle the error and return an error response.
		handleRegistrationError(c, err)
		return
	}

	formatterr := UserResponse.FormatUserResponse(user, "")
	response := formatter.ApiResponse(http.StatusOK, "success", formatterr, "User Registered Successfully")
	c.JSON(http.StatusOK, response)
}

func (h *userController) Login(c *gin.Context) {
	var userInput UserModel.User

	if err := c.ShouldBindJSON(&userInput); err != nil {
		response := formatter.ApiResponse(http.StatusBadRequest, "error", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Cari pengguna berdasarkan alamat email
	var dbUser UserModel.User
	err := h.db.Debug().Where("email = ?", userInput.Email).First(&dbUser).Error
	if err != nil {
		handleLoginError(c, "Something went wrong")
		return
	}

	// Input password dari pengguna
	inputPassword := userInput.Password

	// Password dari database (nilai hash)
	dbPasswordHash := dbUser.Password

	// Perbandingkan password yang di-hash
	if !bcrypt.ComparePassword(inputPassword, dbPasswordHash) {
		handleLoginError(c, "Email or Password is Wrong")
		return
	}

	// Generate token untuk pengguna
	token, err := jwt.GenerateToken(dbUser.ID)
	if err != nil {
		handleLoginError(c, "Something Went Wrong")
		return
	}

	// Format respons
	responseData := UserResponse.FormatUserResponse(dbUser, token)
	response := formatter.ApiResponse(http.StatusOK, "success", responseData, "User Login Successfully")
	c.JSON(http.StatusOK, response)
}

func (h *userController) Update(c *gin.Context) {
	var oldUser UserModel.User
	var newUser UserModel.User

	id := c.Param("userId")

	if err := h.db.First(&oldUser, id).Error; err != nil {
		handleUpdateError(c, err)
		return
	}

	if err := json.NewDecoder(c.Request.Body).Decode(&newUser); err != nil {
		handleUpdateError(c, err)
		return
	}

	if err := h.db.Model(&oldUser).Updates(newUser).Error; err != nil {
		handleUpdateError(c, err)
		return
	}

	response := formatter.ApiResponse(http.StatusOK, "success", nil, "User Updated Successfully")
	c.JSON(http.StatusOK, response)
}

func (h *userController) Delete(c *gin.Context) {
	var user UserModel.User

	id := c.Param("userId")
	if err := h.db.First(&user, id).Error; err != nil {
		handleDeleteError(c, err)
		return
	}

	if err := h.db.Delete(&user).Error; err != nil {
		handleDeleteError(c, err)
		return
	}

	response := formatter.ApiResponse(http.StatusOK, "success", nil, "User Deleted Successfully")
	c.JSON(http.StatusOK, response)
}

func handleRegistrationError(c *gin.Context, err error) {
	errors := formatter.FormatValidationError(err)
	errorMessage := gin.H{"errors": errors}
	response := formatter.ApiResponse(http.StatusUnprocessableEntity, "error", errorMessage, err.Error())
	c.JSON(http.StatusUnprocessableEntity, response)
}

func handleLoginError(c *gin.Context, message string) {
	response := formatter.ApiResponse(http.StatusUnprocessableEntity, "error", nil, message)
	c.JSON(http.StatusUnprocessableEntity, response)
}

func handleUpdateError(c *gin.Context, err error) {
	errors := formatter.FormatValidationError(err)
	errorMessage := gin.H{"errors": errors}
	response := formatter.ApiResponse(http.StatusUnprocessableEntity, "error", errorMessage, err.Error())
	c.JSON(http.StatusUnprocessableEntity, response)
}

func handleDeleteError(c *gin.Context, err error) {
	response := formatter.ApiResponse(http.StatusUnprocessableEntity, "error", nil, err.Error())
	c.JSON(http.StatusUnprocessableEntity, response)
}
