package soap

import (
	"encoding/xml"
)

type Ordertype string

const (
	OrdertypeAsc Ordertype = "asc"

	OrdertypeDesc Ordertype = "desc"
)

type Rowset struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap rowset"`
}

type Row struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap row"`
}

type Column struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap column"`

	Name    string `xml:"name,omitempty"`

	Value   string `xml:"value,omitempty"`
}

type Sort struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap sort"`
	Order   *Ordertype
}

type Sortitem struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap sort_item"`

	Column  string `xml:"column,omitempty"`

	Dir     *Ordertype `xml:"dir,omitempty"`
}

type Filterchildsarr struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap filter_childs_arr"`
}

type Filter struct {
	XMLName  xml.Name `xml:"http://mfisoft.ru/voip/service/soap filter"`

	Type_    string `xml:"type,omitempty"`

	Column   string `xml:"column,omitempty"`

	Operator string `xml:"operator,omitempty"`

	Value    string `xml:"value,omitempty"`

	Childs   *Filterchildsarr //`xml:"childs,omitempty"`
}

