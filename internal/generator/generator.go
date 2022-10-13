package generator

import (
	"model_generator/internal/db"
	"model_generator/internal/language"
)

type mGen struct {
	database db.Database
	language language.Language
}

func New(database db.Database, language language.Language) *mGen {
	return &mGen{database: database, language: language}
}

func (generator *mGen) StartGenFiles(path string) error {
	entities, err := generator.database.GetEntities()
	if err != nil {
		return err
	}
	return generator.language.GenFiles(entities, path)
}
