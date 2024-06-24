// FUTURE IDEA
package sorm

type Attributes []Attribute

func (r Attributes) ColumnNames() []string {
	columnNames := []string{}

	for _, attribute := range r {
		columnNames = append(columnNames, attribute.Name)
	}

	return columnNames
}

func (r Attributes) FieldValues() []any {
	fieldValues := []any{}

	for _, attribute := range r {
		fieldValues = append(fieldValues, attribute.Value)
	}

	return fieldValues
}

type Attribute struct {
	Name       string
	Value      any
	ForeignKey bool
}

func NewAttribute(name string, value any) Attribute {
	return Attribute{
		Name:       name,
		Value:      value,
		ForeignKey: false,
	}
}

func NewForeignKey(name string, value any) Attribute {
	return Attribute{
		Name:       name,
		Value:      value,
		ForeignKey: true,
	}
}

