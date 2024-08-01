package common

type ServerRole struct {
	Name string `json:"name"`
}

func (principal ServerRole) GetName() string {
	return principal.Name
}

func (principal ServerRole) Extract() string {
	return "[DDL_GOES_HERE]" // TODO: extract DDL
}
