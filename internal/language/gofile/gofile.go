package gofile

import "model_generator/internal/models"

type gofile struct{}

func New() *gofile {
	return &gofile{}
}

func GenFiles(entities []models.Entity, path string) error {
	return nil
}
