package db

import "model_generator/internal/models"

// Интерфейс, ограничивающий реализации взаимодействия с БД
type Database interface {
	GetEntities() ([]models.Entity, error)
}
