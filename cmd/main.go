package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	_ "github.com/microsoft/go-mssqldb"

	"github.com/madera-unified/mssql-ddlextractor/internal/config"
	"github.com/madera-unified/mssql-ddlextractor/pkg"
)

func main() {
	configuration := config.GetConfig()

	for _, server := range configuration.Servers {
		hostname := server.Hostname
		port := server.Port
		username := url.QueryEscape(server.Username)
		password := url.QueryEscape(server.Password)
		dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d", username, password, hostname, port)
		log.Printf("Connecting to %s:%d\n", hostname, port)

		db, err := sql.Open("sqlserver", dsn)

		if err != nil {
			log.Fatalf("Error creating connection pool:\n%s\n", err.Error())
			continue
		}

		err = db.Ping()
		if err != nil {
			log.Fatalf("Error opening connection:\n%s\n", err.Error())
		}

		log.Printf("Connected to %s:%d\n", hostname, port)
		pkg.Init(db)
	}
}
