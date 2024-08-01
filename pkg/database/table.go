package database

type Table struct {
	Name        string       `json:"name"`
	Columns     []Column     `json:"columns"`
	Keys        []Key        `json:"keys"`
	Constraints []Constraint `json:"constraints"`
	Indexes     []Index      `json:"indexes"`
}

func (table Table) Extract() string {
	return "[DDL_GOES_HERE]" // TODO: extract DDL
}
