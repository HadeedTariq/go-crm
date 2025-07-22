package leads

import (
	"github.com/gin-gonic/gin"
)

type ContactData struct {
	FirstName   string `json:"name" validate:"required,min=2,max=100"`
	LastName    string `json:"value" validate:""`
	Email       string `json:"stage" validate:"required,dealstage"`
	Phone       string `json:"stage" validate:"required,dealstage"`
	CompanyName string `json:"stage" validate:"required,dealstage"`
	LeadStatus  string `json:"lead_status" validate:"required,leadstatus"`
	Source      string `json:"stage" validate:"required,dealstage"`
}

func CreateContact(c *gin.Context) {

	// ~ so now the validation part is done

}
