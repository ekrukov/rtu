package soap

import (
	"encoding/xml"
)

type Ordertype string

const (
	OrdertypeAsc Ordertype = "asc"
	OrdertypeDesc Ordertype = "desc"
)


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
	XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap p_sort"`
	Items   []Sortitem `xml:"http://mfisoft.ru/voip/service/soap item"`
}

type Sortitem struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap sort_item"`
	Column  string `xml:"column,omitempty"`
	Dir     Ordertype `xml:"dir,omitempty"`
}
/*
type Filterchildsarr struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap filter_childs_arr"`
	Filters []Filter `xml:"http://mfisoft.ru/voip/service/soap item"`
}*/
// TODO only simple filter

type Filter struct {
	XMLName  xml.Name `xml:"http://mfisoft.ru/voip/service/soap p_filter"`
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
