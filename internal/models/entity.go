package models

// Структура, описывающая сущность БД: таблица/представление
type Entity struct {
	SchemaName string
	EntityName string
	Comment    string
	Properties []Property
}
