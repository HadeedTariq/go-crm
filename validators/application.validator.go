package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func InitValidator() {
	Validate = validator.New()

	Validate.RegisterValidation("userrole", func(fl validator.FieldLevel) bool {
		role := strings.ToLower(fl.Field().String())

		validRoles := map[string]bool{
			"sales_rep":            true,
			"marketing_manager":    true,
			"customer_service_rep": true,
			"admin":                true,
		}

		return validRoles[role]
	})
	Validate.RegisterValidation("dealstage", func(fl validator.FieldLevel) bool {
		stage := strings.ToLower(fl.Field().String())

		validStages := map[string]bool{
			"prospecting": true,
			"negotiation": true,
			"closed_won":  true,
			"closed_lost": true,
		}

		return validStages[stage]
	})
}
