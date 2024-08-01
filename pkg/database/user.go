package database

import "github.com/madera-unified/mssql-ddlextractor/pkg/common"

type User struct {
	Name          string       `json:"name"`
	Login         common.Login `json:"login"`
	DefaultSchema Schema       `json:"defaultSchema"`
}

func (user User) Extract() string {
	return "[DDL_GOES_HERE]" // TODO: extract DDL
}
