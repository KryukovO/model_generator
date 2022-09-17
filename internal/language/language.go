package language

import "model_generator/internal/models"

type Language interface {
	GenFiles([]models.Entity, string) error
}
