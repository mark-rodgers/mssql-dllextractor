package role

type ApplicationRole struct {
	Name string `json:"name"`
}

func (principal ApplicationRole) GetName() string {
	return principal.Name
}

func (principal ApplicationRole) Extract() string {
	return "[DDL_GOES_HERE]" // TODO: extract DDL
}
