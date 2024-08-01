package database

type Key struct {
	Name      string `json:"name"`
	DataType  string `json:"dataType"`
	AllowNull bool   `json:"allowNull"`
}

func (key Key) Extract() string {
	return "[DDL_GOES_HERE]" // TODO: extract DDL
}
