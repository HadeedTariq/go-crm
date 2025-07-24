package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type UserRole string

const (
	RoleSalesRep           UserRole = "sales_rep"
	RoleMarketingManager   UserRole = "marketing_manager"
	RoleCustomerServiceRep UserRole = "customer_service_rep"
	RoleAdmin              UserRole = "admin"
)

func (ur UserRole) IsValid() bool {
	switch ur {
	case RoleSalesRep,
		RoleMarketingManager,
		RoleCustomerServiceRep,
		RoleAdmin:
		return true
	}
	return false
}

func validateUserRole(fl validator.FieldLevel) bool {
	role := UserRole(strings.ToLower(fl.Field().String()))

	return role.IsValid()
}
