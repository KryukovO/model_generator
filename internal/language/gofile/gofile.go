package gofile

import (
	"fmt"
	"log"
	"model_generator/internal/models"
	"os"
	"strings"
)

type gofile struct{}

func New() *gofile {
	return &gofile{}
}

// Генерация файлов .go для таблиц и представлений БД по указанному пути
func (gofile) GenFiles(entities []models.Entity, path string) error {
	// Удаляем слеш в конце пути
	if path[len([]rune(path))-1] == '/' {
		path = path[:len([]rune(path))-1]
	}

	// Удаление существующей папки пути
	err := os.RemoveAll(path)
	if err != nil {
		log.Fatal(err)
	}

	// Создание папки заново
	err = os.MkdirAll(path, 0755)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Directory %v recreated\n", path)

	// Генерация файлов
	for _, entity := range entities {
		// Создаём файл модели
		filename := fmt.Sprintf("%v-%v.go", entity.SchemaName, entity.EntityName)
		f, err := os.Create(fmt.Sprintf("%v/%v", path, filename))
		if err != nil {
			return err
		}
		defer f.Close()

		log.Printf("File %v/%v created\n", path, filename)

		// Записываем данные в файл
		pathParts := strings.Split(path, "/")
		fileContent := fmt.Sprintf("package %v\n\n// %s\ntype %s struct {\n", pathParts[len(pathParts)-1], entity.Comment, strings.Title(entity.EntityName))
		for _, property := range entity.Properties {
			// TODO: подумать над преобразованием к конкретному типу, а пока пусть все будет string, []string и bool
			fieldType := ""
			switch {
			case property.DataType == "boolean":
				fieldType = "bool"
			case property.DataType == "[]boolean":
				fieldType = "[]bool"
			case property.DataType[0] == '[':
				fieldType = "[]string"
			default:
				fieldType = "string"
			}
			fileContent += fmt.Sprintf("\t%v %v `json:\"%v\"` // %v\n", strings.Title(property.ColumnName), fieldType, property.ColumnName, property.Comment)
		}
		fileContent += "}"

		_, err = f.WriteString(fileContent)
		if err != nil {
			return err
		}
	}

	return nil
}
