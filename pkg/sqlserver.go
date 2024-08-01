package pkg

import "github.com/madera-unified/mssql-ddlextractor/pkg/common"

type SQLServer struct {
	Databases []Database `json:"databases"`
	Security  struct {
		Logins []common.Login `json:"logins"`
	} `json:"security"`
}
