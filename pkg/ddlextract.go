package pkg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/madera-unified/mssql-ddlextractor/pkg/common"
	"github.com/madera-unified/mssql-ddlextractor/pkg/database"
)

var DB *sql.DB
var Infrastructure *SQLServer

func Init(db *sql.DB) {
	DB = db
	Infrastructure = buildInfrastructure()

	json, err := json.MarshalIndent(Infrastructure, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}
	log.Println(string(json))
}

func buildInfrastructure() *SQLServer {
	var infrastructure SQLServer

	infrastructure.Security.Logins = getServerLogins()
	infrastructure.Databases = getDatabases()

	return &infrastructure
}

func getSqlFile(filename string) string {
	sqlFile := fmt.Sprintf("sql/%s", filename)
	sql, err := ioutil.ReadFile(sqlFile)
	if err != nil {
		log.Fatalf("Error reading SQL file: %v", err)
	}
	return string(sql)
}

func getServerLogins() []common.Login {
	return []common.Login{}
}

func getDatabases() []Database {
	exclusions := []string{"'master'", "'tempdb'", "'model'", "'msdb'", "'distribution'", "'SSISDB'"}

	query := fmt.Sprintf(getSqlFile("get_databases.sql"), strings.Join(exclusions, ", "))
	rows, err := DB.Query(query)
	if err != nil {
		log.Fatalf("Error selecting databases:\n%s\n", err.Error())
	}
	defer rows.Close()

	var dbs []Database
	for rows.Next() {
		var db Database

		if err := rows.Scan(&db.Name, &db.Owner); err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}

		getTables(&db)

		dbs = append(dbs, db)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		log.Fatalf("Error during iteration: %v", err)
	}

	return dbs
}

func getTables(db *Database) {
	query := fmt.Sprintf(getSqlFile("get_tables_and_views.sql"), db.Name)
	rows, err := DB.Query(query)
	if err != nil {
		log.Fatalf("Error selecting tables/views for %s:\n%s\n", db.Name, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var tableSchema string
		var tableName string
		var tableType string

		if err := rows.Scan(&tableSchema, &tableName, &tableType); err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}

		switch tableType {
		case "BASE TABLE":
			db.Tables = append(db.Tables, database.Table{
				Name: tableName,
			})
		case "VIEW":
			db.Views = append(db.Views, database.View{
				Name: tableName,
			})
		default:
			log.Printf("Unknown table type: %s\n", tableType)
		}
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		log.Fatalf("Error during iteration: %v", err)
	}
}
