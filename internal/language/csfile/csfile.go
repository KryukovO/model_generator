package csfile

import (
	"errors"
	"model_generator/internal/models"
)

type csfile struct{}

func New() *csfile {
	return &csfile{}
}

func (csfile) GenFiles(entities []models.Entity, path string) error {
	return errors.New("генерация моделей для языка C# пока не реализована")
}
