package role

type DatabaseRole struct {
	Name string `json:"name"`
}

func (principal DatabaseRole) GetName() string {
	return principal.Name
}

func (principal DatabaseRole) Extract() string {
	return "[DDL_GOES_HERE]" // TODO: extract DDL
}
