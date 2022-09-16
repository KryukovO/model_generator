package pgsql

import (
	"database/sql"
	"fmt"
	"model_generator/internal/config"
	"model_generator/internal/models"

	_ "github.com/lib/pq"
)

// Структура для взаимодействия с БД
type PgsqlDB struct {
	connectionStr string
}

// Конструктор PgsqlDB
func New(conf *config.Config) (*PgsqlDB, error) {
	pgsqlDB := &PgsqlDB{
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			conf.Host,
			conf.Port,
			conf.User,
			conf.Password,
			conf.DataBase,
		),
	}

	db, err := sql.Open("postgres", pgsqlDB.connectionStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return pgsqlDB, nil
}

func (pgsqlDB *PgsqlDB) GetEntities() ([]models.Entity, error) {
	db, err := sql.Open("postgres", pgsqlDB.connectionStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	entityQuery := `SELECT 
						ent.oid AS entity_oid,
						nsp.nspname AS schema_name,
						ent.relname AS entity_name,
						COALESCE(obj_description(ent.oid), format('%s.%s', nsp.nspname, ent.relname)) AS "comment"
					FROM pg_class ent
					JOIN pg_namespace nsp ON nsp.oid = ent.relnamespace
					WHERE (ent.relkind = 'r' OR ent.relkind = 'v')
						AND nsp.nspname <> ALL (ARRAY['pg_catalog'::text, 'information_schema'::text]) AND nsp.nspname NOT LIKE 'pg_toast%';`

	propertyQuery := `SELECT 
						att.attname AS column_name,
						format_type(att.atttypid, NULL)::information_schema.character_data AS data_type,
						COALESCE(d.description, att.attname::TEXT) AS "comment"
					FROM pg_attribute att
					LEFT JOIN pg_description d ON d.objsubid = att.attnum AND d.objoid = att.attrelid
					WHERE att.attrelid = $1 AND att.attnum > 0 AND NOT att.attisdropped;`

	entities := make([]models.Entity, 0)
	entityRows, err := db.Query(entityQuery)
	if err != nil {
		return nil, err
	}

	for entityRows.Next() {
		entity := models.Entity{}
		var entityOid string

		err = entityRows.Scan(&entityOid, &entity.SchemaName, &entity.EntityName, &entity.Comment)
		if err != nil {
			return nil, err
		}

		properties := make([]models.Property, 0)
		propertyRow, err := db.Query(propertyQuery, entityOid)
		if err != nil {
			return nil, err
		}

		for propertyRow.Next() {
			property := models.Property{}
			err = propertyRow.Scan(&property.ColumnName, &property.DataType, &property.Comment)
			if err != nil {
				return nil, err
			}

			properties = append(properties, property)
		}
		entity.Properties = properties

		entities = append(entities, entity)
	}

	return entities, nil
}
