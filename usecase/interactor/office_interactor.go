package interactor

import (
	"somer/domain"
	"somer/usecase/presenter"
	"somer/usecase/repository"
)

type officeInteractor struct {
	OfficeRepository repository.OfficeRepository
	OfficePresenter  presenter.OfficePresenter
}

// OfficeInteractor wraps office interactor method
type OfficeInteractor interface {
	Fetch() ([]domain.Office, error)
}

// NewOfficeInteractor initialize office interactor
func NewOfficeInteractor(r repository.OfficeRepository, p presenter.OfficePresenter) OfficeInteractor {
	return &officeInteractor{r, p}
}

func (of *officeInteractor) Fetch() ([]domain.Office, error) {
	o, err := of.OfficeRepository.Fetch()
	if err != nil {
		return nil, err
	}
	return of.OfficePresenter.ResponseOffices(o), nil
}
