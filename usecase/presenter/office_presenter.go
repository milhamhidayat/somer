package presenter

import "somer/domain"

// OfficePresenter wraps office presenter method
type OfficePresenter interface {
	ResponseOffices(offices []domain.Office) []domain.Office
}
