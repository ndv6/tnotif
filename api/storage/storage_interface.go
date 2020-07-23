package storage

import (
	"github.com/ndv6/tnotif/models"
)

type Storage interface {
	Create(obj models.LogMail) error
	List() ([]models.LogMail, error)
}
