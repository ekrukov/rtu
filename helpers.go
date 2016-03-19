package rtu

import (
	"errors"
	"strings"
)

var (
	errFilterUnknownField = errors.New("Unknown field in filter map")
	errFilterUnknownCondition = errors.New("Unknown condition in filter string")
)

func stringToFilter(s string) (f *requestFilter, err error) {
	f = new(requestFilter)
	f.P_filter.Type_ = "cond"
	for _, cond := range filterConditions {
		if condPosition := strings.Index(s, cond); condPosition != -1 {
			valuePos := len(cond) + condPosition
			column := strings.TrimSpace(s[:condPosition])
			value := strings.TrimSpace(s[valuePos:])
			if cond == "like" || cond == "not like" {
				value = "%" + "%"
			}
			f.P_filter.Column = column
			f.P_filter.Operator = cond
			f.P_filter.Value = value
			return f, nil
		}
	}
	return f, errFilterUnknownCondition
}

func mapToSort(m map[string]Ordertype) (s *requestSort, err error) {
	items := []requestSortItem{}
	for column, dir := range m {
		items = append(items, requestSortItem{
			Column: column,
			Dir: dir,
		})
	}
	s = new(requestSort)
	s.P_sort.Items = items
	return s, err
}

func mapToFilter(m map[string]string) (f *requestFilter, err error) {
	f = new(requestFilter)
	f.P_filter.Type_, err = checkInMap(m, "type")
	f.P_filter.Column, err = checkInMap(m, "column")
	f.P_filter.Operator, err = checkInMap(m, "operator")
	f.P_filter.Value, err = checkInMap(m, "value")
	return f, err
}

func checkInMap(m map[string]string, key string) (value string, err error) {
	if value, ok := m[key]; ok {
		return value, nil
	}
	return "", errFilterUnknownField
}

func mapToRow(m map[string]string) (r *requestRow, err error) {
	columnMap := []requestColumn{}
	for key, value := range m {
		columnMap = append(columnMap, requestColumn{
			Name: key,
			Value: value,
		})
	}
	r = &requestRow{Items: columnMap}
	return r, err
}

func mapsToRowset(m []map[string]string) (r *requestRowset, err error) {
	rows := []requestRow{}
	for _, row := range m {
		rowMap, err := mapToRow(row)
		rows = append(rows, *rowMap)
		if err != nil {
			return nil, err
		}
	}
	r = new(requestRowset)
	r.P_rowset.Rows = rows
	return r, err
}


