package repository

import (
	"somer/domain"
)

// OfficeRepository wraps office repository operation
type OfficeRepository interface {
	Fetch() ([]domain.Office, error)
}
