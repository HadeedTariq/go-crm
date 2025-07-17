package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hadeedtariq/go-crm/config"
	"github.com/hadeedtariq/go-crm/models"
	"github.com/hadeedtariq/go-crm/validators"
	"golang.org/x/crypto/bcrypt"
)

type RegisterData struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Role     string `json:"role" validate:"required,userrole"`
	Password string `json:"password" validate:"required,min=6"`
}

func RegisterUser(c *gin.Context) {
	var data RegisterData

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format", "details": err.Error()})
		return
	}

	if err := validators.Validate.Struct(data); err != nil {
		var validationErrors []string
		for _, e := range err.(validator.ValidationErrors) {
			msg := fmt.Sprintf("Validation failed on '%s': rule '%s'", e.Field(), e.Tag())
			validationErrors = append(validationErrors, msg)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Validation failed",
			"details": validationErrors,
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password", "details": err.Error()})
		return
	}

	user := models.User{
		Name:     data.Name,
		Email:    data.Email,
		Role:     data.Role,
		Password: string(hashedPassword),
	}

	if result := config.DB.Create(&user); result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to create user",
				"details": result.Error.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
