package database

type Column struct {
	Name      string `json:"name"`
	DataType  string `json:"dataType"`
	AllowNull bool   `json:"allowNull"`
}

func (column Column) Extract() string {
	return "[DDL_GOES_HERE]" // TODO: extract DDL
}
