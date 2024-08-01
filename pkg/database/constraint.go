package database

type Constraint struct {
	Name      string `json:"name"`
	DataType  string `json:"dataType"`
	AllowNull bool   `json:"allowNull"`
}

func (constraint Constraint) Extract() string {
	return "[DDL_GOES_HERE]" // TODO: extract DDL
}
