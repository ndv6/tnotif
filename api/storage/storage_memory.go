package storage

import (
	"github.com/ndv6/tnotif/models"
)

type memory struct {
	store map[int32]models.LogMail
}

func newMemory() memory {
	return memory{
		store: make(map[int32]models.LogMail),
	}
}
func (o memory) Create(obj models.LogMail) error {
	o.store[int32(len(o.store))] = obj
	return nil
}

func (o memory) List() ([]models.LogMail, error) {
	result := []models.LogMail{}
	for _, v := range o.store {
		result = append(result, v)
	}
	return result, nil
}
