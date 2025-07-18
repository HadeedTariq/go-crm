package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hadeedtariq/go-crm/config"
	"github.com/hadeedtariq/go-crm/models"
	"github.com/hadeedtariq/go-crm/utils"
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

type LoginData struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type MyCustomClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}

var jwtSecret = utils.GetEnv("JWT_SECRET_KEY", "hadeedbhai")

func LoginUser(c *gin.Context) {
	var data LoginData

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input format",
			"details": err.Error(),
		})
		return
	}

	var user models.User
	result := config.DB.Where("email = ?", data.Email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	claims := MyCustomClaims{
		UserID: user.ID.String(),
		Role:   user.Role,
		Email:  user.Email,
		Name:   user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "Go-Crm",
			Subject:   user.ID.String(),
		},
	}

	// Use HS256 for string-based secret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}

	// Cookie settings – adjust domain in prod
	c.SetCookie(
		"jwt_token",
		tokenString,
		int(24*time.Hour.Seconds()),
		"/",
		"localhost",
		true, // Secure (HTTPS only)
		true, // HttpOnly
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
	})
}
