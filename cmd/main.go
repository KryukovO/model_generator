package main

import (
	"encoding/json"
	"flag"
	"log"
	"model_generator/internal/config"
	"model_generator/internal/db/pgsql"
	"model_generator/internal/generator"
	"model_generator/internal/language"
	"model_generator/internal/language/csfile"
	"model_generator/internal/language/gofile"
	"os"
)

const (
	defaultConfig   string = "configs/default.json"
	defaultLanguage string = "go"
	defaultOutput   string = "models/"
)

var (
	configPath string
	lang       string
	output     string
)

func init() {
	flag.StringVar(&configPath, "config", defaultConfig, "config path")
	flag.StringVar(&lang, "lang", defaultLanguage, "language for which models are generated")
	flag.StringVar(&output, "o", defaultOutput, "path to output directory")
}

func main() {
	// Парсим флаги
	flag.Parse()

	// Считывание конфига
	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	var cnfg config.Config
	err = json.Unmarshal(data, &cnfg)
	if err != nil {
		log.Fatal(err)
	}

	// Подключение к БД
	db, err := pgsql.New(&cnfg)
	if err != nil {
		log.Fatal(err)
	}

	// Обработка языков
	var genLang language.Language
	switch lang {
	case "go":
		log.Println(".go files will be generated")
		genLang = gofile.New()
	case "cs":
		log.Println(".cs files will be generated")
		genLang = csfile.New()
	default:
		log.Fatal("Chosen language does not support")
	}

	// Генерация моделей
	gen := generator.New(db, genLang)
	err = gen.StartGenFiles(output)
	if err != nil {
		log.Fatal(err)
	}
}
