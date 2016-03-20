package rtu

import (
	"errors"
	"strings"
)

var (
	errFilterUnknownCondition = errors.New("Unknown condition in filter string")
)

func stringToFilter(s string) (f *requestFilterItem, err error) {
	f = new(requestFilterItem)
	f.Type_ = simpleFilter
	for _, cond := range filterConditions {
		if condPosition := strings.Index(s, cond); condPosition != -1 {
			valuePos := len(cond) + condPosition
			column := strings.TrimSpace(s[:condPosition])
			value := strings.TrimSpace(s[valuePos:])
			if cond == "like" || cond == "not like" {
				value = "%" + value + "%"
			}
			f.Column = column
			f.Operator = cond
			f.Value = value
			return f, nil
		}
	}
	return f, errFilterUnknownCondition
}

func mapToSort(m map[string]OrderType) (s *requestSort, err error) {
	items := []*requestSortItem{}
	for column, dir := range m {
		items = append(items, &requestSortItem{
			Column: column,
			Dir: dir,
		})
	}
	s = new(requestSort)
	s.P_sort.Items = items
	return s, err
}

func sliceToChildFilters(s []string) (fi []*requestFilterItem, err error) {
	fi = []*requestFilterItem{}
	for _, sf := range s {
		cf, err := stringToFilter(sf)
		if err != nil {
			return nil, err
		}
		fi = append(fi, cf)
	}
	return fi, nil
}

func mapToRow(m map[string]string) (r *requestRow, err error) {
	columnMap := []*requestColumn{}
	for key, value := range m {
		columnMap = append(columnMap, &requestColumn{
			Name: key,
			Value: value,
		})
	}
	r = &requestRow{Items: columnMap}
	return r, err
}

func mapsToRowset(m []map[string]string) (r *requestRowset, err error) {
	rows := []*requestRow{}
	for _, row := range m {
		rowMap, err := mapToRow(row)
		rows = append(rows, rowMap)
		if err != nil {
			return nil, err
		}
	}
	r = new(requestRowset)
	r.P_rowset.Rows = rows
	return r, err
}


