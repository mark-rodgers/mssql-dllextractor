package pkg

import (
	"github.com/madera-unified/mssql-ddlextractor/pkg/database"
)

type Database struct {
	Name   string           `json:"name"`
	Owner  string           `json:"owner"` // TODO: change to common.Login
	Tables []database.Table `json:"tables"`
	Views  []database.View  `json:"views"`
}

func (database Database) Extract() string {
	return "[DDL_GOES_HERE]" // TODO: extract DDL
}
