package leads

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hadeedtariq/go-crm/config"
	"github.com/hadeedtariq/go-crm/models"
	"github.com/hadeedtariq/go-crm/validators"
)

type ContactData struct {
	FirstName   string `json:"first_name" validate:"required,min=2,max=100"`
	LastName    string `json:"last_name"`                                  // No specific validation given for now, remove 'validate:""'
	Email       string `json:"email" validate:"required,email"`            // Correct tag for email, use 'email' validator
	Phone       string `json:"phone"`                                      // Correct tag for phone, remove 'validate:""' if no specific validation
	CompanyName string `json:"company_name"`                               // Correct tag for company name, remove 'validate:""' if no specific validation
	LeadStatus  string `json:"lead_status" validate:"required,leadstatus"` // Correct tag and custom validator
	Source      string `json:"source" validate:"required,leadsource"`      // Correct tag and custom validator (e.g., 'sourceval' to avoid conflict)
}

func CreateContact(c *gin.Context) {
	var data ContactData

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

	contact := models.Contact{
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		Email:       data.Email,
		Phone:       data.Phone,
		CompanyName: data.CompanyName,
		LeadStatus:  data.LeadStatus,
		Source:      data.Source,
	}

	if result := config.DB.Create(&contact); result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, gin.H{"error": "Contact with this Email already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to create contact",
				"details": result.Error.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Contact created successfully"})

}
