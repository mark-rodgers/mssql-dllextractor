package common

type Login struct {
	Name     string `json:"name"`
	AuthType string `json:"authType"` // TODO: constrain to SQL_LOGIN || WINDOWS_LOGIN
}

func (login Login) Extract() string {
	return "[DDL_GOES_HERE]" // TODO: extract DDL
}
