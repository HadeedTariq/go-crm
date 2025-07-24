package validators

import (
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func InitValidator() {
	Validate = validator.New()

	Validate.RegisterValidation("userrole", validateUserRole)
	Validate.RegisterValidation("dealstage", validateDealStage)
	Validate.RegisterValidation("leadstatus", validateLeadStatus)
	Validate.RegisterValidation("leadsource", validateLeadSource)
}
