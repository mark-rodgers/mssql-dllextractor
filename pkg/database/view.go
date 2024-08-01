package database

type View struct {
	Name string `json:"name"`
}

func (view View) Extract() string {
	return "[DDL_GOES_HERE]" // TODO: extract DDL
}
