package helper

import "strings"

type FilterBuilder struct {
	conditions []string
	args       []interface{}
}

func NewFilterBuilder() *FilterBuilder {
	return &FilterBuilder{
		conditions: []string{},
		args:       []interface{}{},
	}
}

// untuk value optional (pointer)
func (f *FilterBuilder) Add(condition string, value interface{}) {
	if value == nil {
		return
	}

	f.conditions = append(f.conditions, condition)

	switch v := value.(type) {
	case *int:
		f.args = append(f.args, *v)
	case *string:
		f.args = append(f.args, *v)
	default:
		f.args = append(f.args, v)
	}
}

// untuk kondisi custom (BETWEEN, dll)
func (f *FilterBuilder) AddRaw(condition string, values ...interface{}) {
	f.conditions = append(f.conditions, condition)
	f.args = append(f.args, values...)
}

// untuk IN query
func (f *FilterBuilder) AddIn(column string, values []int) {
	if len(values) == 0 {
		return
	}

	placeholders := strings.Repeat("?,", len(values))
	placeholders = placeholders[:len(placeholders)-1]

	f.conditions = append(f.conditions, column+" IN ("+placeholders+")")

	for _, v := range values {
		f.args = append(f.args, v)
	}
}

func (f *FilterBuilder) BuildWhere() (string, []interface{}) {
	if len(f.conditions) == 0 {
		return "", f.args
	}

	return " WHERE " + strings.Join(f.conditions, " AND "), f.args
}
