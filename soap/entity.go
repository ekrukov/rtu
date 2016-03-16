package soap

import (
	"encoding/xml"
	"errors"
)

type Ordertype string

const (
	OrdertypeAsc Ordertype = "asc"

	OrdertypeDesc Ordertype = "desc"
)

func MapToFilter(m map[string]string) (f *Filter, e error) {
	f = new(Filter)
	f.Type_, e = checkInMap(m, "type")
	f.Column, e = checkInMap(m, "column")
	f.Operator, e = checkInMap(m, "operator")
	f.Value, e = checkInMap(m, "value")
	return f, e
}

func checkInMap(m map[string]string, key string) (value string, err error) {
	if value, ok := m[key]; ok {
		return value, nil
	}
	return "", errors.New("Unknown field in filter map")
}

func MapToRow(m *map[string]string) (r *Row, e error) {
	columnMap := []Column{}
	for key, value := range *m {
		columnMap = append(columnMap, Column{
			Name: key,
			Value: value,
		})
	}
	r = &Row{Items: columnMap}
	return r, e
}

func MapsToRowset(m *[]map[string]string) (r *Rowset, e error) {
	rows := []Row{}
	for _, row := range *m {
		rowMap, e := MapToRow(&row)
		rows = append(rows, *rowMap)
		if e != nil {
			return nil, e
		}
	}
	r = &Rowset{Rows: rows}
	return r, e
}

// TODO XML attr not in use now


type Rowset struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap rowset"`
	Rows    []Row `xml:"item"`
}

type Row struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap row"`
	Items   []Column `xml:"http://mfisoft.ru/voip/service/soap item"`
}

type Column struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap column"`
	Name    string `xml:"name,omitempty"`
	Value   string `xml:"value,omitempty"`
}

type Sort struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap sort"`
	Item    Sortitem `xml:"http://mfisoft.ru/voip/service/soap item"`
}

type Sortitem struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap sort_item"`
	Column  string `xml:"column,omitempty"`
	Dir     *Ordertype `xml:"dir,omitempty"`
}

type Filterchildsarr struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap filter_childs_arr"`
	Filters []Filter `xml:"http://mfisoft.ru/voip/service/soap item"`
}

// TODO only simple filter

type Filter struct {
	XMLName  xml.Name `xml:"http://mfisoft.ru/voip/service/soap filter"`
	Type_    string `xml:"type,omitempty"`
	Column   string `xml:"column,omitempty"`
	Operator string `xml:"operator,omitempty"`
	Value    string `xml:"value,omitempty"`
	//Childs   Filterchildsarr `xml:"childs,omitempty"`
}

type ResponceRowset struct {
	XMLName xml.Name `xml:"result"`
	Rows    []ResponceRow `xml:"item"`
}

type ResponceRow struct {
	XMLName xml.Name `xml:"item"`
	Items   []ResponceColumn `xml:"item"`
}

type ResponceColumn struct {
	XMLName xml.Name `xml:"item"`
	Key     string `xml:"key"`
	Value   string `xml:"value"`
}
