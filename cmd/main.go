package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"model_generator/internal/config"
	"os"
)

const (
	defaultConfig   string = "configs/default.json"
	defaultLanguage string = "go"
)

var (
	configPath string
	language   string
)

func init() {
	flag.StringVar(&configPath, "config", defaultConfig, "config path")
	flag.StringVar(&language, "lang", defaultLanguage, "language for which models are generated")
}

func main() {
	flag.Parse()

	switch language {
	case "go":
		println(".go files generated")
	case "cs":
		println(".cs files generated")
	default:
		log.Fatal("Language does not support")
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))

	var cnfg config.Config
	err = json.Unmarshal(data, &cnfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cnfg)
}
