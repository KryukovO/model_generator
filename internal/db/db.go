package db

import "model_generator/internal/models"

// Интерфейс, ограничивающий реализации взаимодействия с БД
type DataBase interface {
	GetEntities() ([]models.Entity, error)
}
