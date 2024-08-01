package database

type Index struct {
	Name      string `json:"name"`
	DataType  string `json:"dataType"`
	AllowNull bool   `json:"allowNull"`
}

func (index Index) Extract() string {
	return "[DDL_GOES_HERE]" // TODO: extract DDL
}
