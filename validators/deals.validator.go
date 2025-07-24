package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type DealStage string

const (
	DealStageProspecting DealStage = "prospecting"
	DealStageNegotiation DealStage = "negotiation"
	DealStageClosedWon   DealStage = "closed_won"
	DealStageClosedLost  DealStage = "closed_lost"
)

func (ds DealStage) IsValid() bool {
	switch ds {
	case DealStageProspecting,
		DealStageNegotiation,
		DealStageClosedWon,
		DealStageClosedLost:
		return true
	}
	return false
}

func validateDealStage(fl validator.FieldLevel) bool {

	stage := DealStage(strings.ToLower(fl.Field().String()))

	return stage.IsValid()
}
