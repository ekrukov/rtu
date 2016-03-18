package rtu

import (
	"errors"
)

func MapToSort(m map[string]Ordertype) (s *Sort, err error) {
	items := []Sortitem{}
	for column, dir := range m {
		items = append(items, Sortitem{
			Column: column,
			Dir: dir,
		})
	}
	s = &Sort{Items: items}
	return s, err
}

func MapToFilter(m map[string]string) (f *Filter, err error) {
	f = new(Filter)
	f.Type_, err = checkInMap(m, "type")
	f.Column, err = checkInMap(m, "column")
	f.Operator, err = checkInMap(m, "operator")
	f.Value, err = checkInMap(m, "value")
	return f, err
}

func checkInMap(m map[string]string, key string) (value string, err error) {
	if value, ok := m[key]; ok {
		return value, nil
	}
	return "", errors.New("Unknown field in filter map")
}

func MapToRow(m map[string]string) (r *Row, err error) {
	columnMap := []Column{}
	for key, value := range m {
		columnMap = append(columnMap, Column{
			Name: key,
			Value: value,
		})
	}
	r = &Row{Items: columnMap}
	return r, err
}

func MapsToRowset(m []map[string]string) (r *Rowset, err error) {
	rows := []Row{}
	for _, row := range m {
		rowMap, err := MapToRow(row)
		rows = append(rows, *rowMap)
		if err != nil {
			return nil, err
		}
	}
	r = &Rowset{Rows: rows}
	return r, err
}


