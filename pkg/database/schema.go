package database

import "github.com/madera-unified/mssql-ddlextractor/pkg/common"

type Schema struct {
	Name  string           `json:"name"`
	Owner common.Principal `json:"owner"`
}

func (schema Schema) Extract() string {
	return "[DDL_GOES_HERE]" // TODO: extract DDL
}
