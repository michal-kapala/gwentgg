package services

import (
	"gorm.io/gorm"
)

type Model interface {
	ToModel() (gorm.Model, error)
}

// Model with a foreign key.
type Submodel interface {
	ToModel(fkey string) (gorm.Model, error)
}

type LocalizationInfo struct {
	Type   string `json:"type"`
	Source string `json:"source"`
	Value  string `json:"value"`
}

type ImageAsset struct {
	ID string `json:"id"`
	Type string `json:"type"`
	Source *string `json:"source"`
}
