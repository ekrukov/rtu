package soap

import (
	"encoding/xml"
)

type Methodtype string

const (
	SelectMethod Methodtype = "SelectRowset"
	InsertMethod Methodtype = "InsertRowset"
	UpdateMethod Methodtype = "UpdateRowset"
	DeleteMethod Methodtype = "DeleteRowset"
	CountMethod Methodtype = "CountRowset"
	DescribeMethod Methodtype = "DescribeColumns"
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

type SelectRowsetRequest struct {
	XMLName    xml.Name `xml:"http://mfisoft.ru/voip/service/soap selectRowset"`
	P_table_hi string `xml:"p_table_hi"`
	P_filter   Filter `xml:"http://mfisoft.ru/voip/service/soap p_filter,omitempty"`
	P_sort     Sort `xml:"http://mfisoft.ru/voip/service/soap p_sort,omitempty"`
	P_limit    int `xml:"p_limit"`
	P_offset   int `xml:"p_offset"`
}

type SelectRowsetResponce struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap selectRowsetResponse"`
	Result  ResponceRowset `xml:"result"`
}

/**
 *	Insert Request structures
 */

type InsertRowsetRequest struct {
	XMLName    xml.Name `xml:"http://mfisoft.ru/voip/service/soap insertRowset"`
	P_table_hi string `xml:"p_table_hi"`
	P_rowset   Rowset `xml:"p_rowset"`
}

type InsertRowsetResponce struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap insertRowsetResponse"`
	Result  int `xml:"result"`
}

/**
 *	Update Request structures
 */

type UpdateRowsetRequest struct {
	XMLName    xml.Name `xml:"http://mfisoft.ru/voip/service/soap updateRowset"`
	P_table_hi string `xml:"p_table_hi"`
	P_rowset   Rowset `xml:"p_rowset"`
	Filter     Filter `xml:"http://mfisoft.ru/voip/service/soap p_filter"`
}

type UpdateRowsetResponce struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap updateRowsetResponse"`
	Result  int `xml:"result"`
}
/**
 *	Delete Request structures
 */

type DeleteRowsetRequest struct {
	XMLName    xml.Name `xml:"http://mfisoft.ru/voip/service/soap deleteRowset"`
	P_table_hi string `xml:"p_table_hi"`
	P_rowset   Rowset `xml:"p_rowset"`
	Filter     Filter `xml:"http://mfisoft.ru/voip/service/soap p_filter"`
}

type DeleteRowsetResponce struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap deleteRowsetResponse"`
	Result  int `xml:"result"`
}

/**
 *	Count Request structures
 */

type CountRowsetRequest struct {
	XMLName    xml.Name `xml:"http://mfisoft.ru/voip/service/soap countRowset"`
	P_table_hi string `xml:"p_table_hi"`
	Filter     Filter `xml:"http://mfisoft.ru/voip/service/soap p_filter,omitempty"`
}

type CountRowsetResponce struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap countRowsetResponse"`
	Result  int `xml:"result"`
}

/**
 *	Describe Request structures
 */

type DescribeColumnsRequest struct {
	XMLName    xml.Name `xml:"http://mfisoft.ru/voip/service/soap describeColumns"`
	P_table_hi string `xml:"p_table_hi"`
}

type DescribeColumnsResponce struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap describeColumnsResponse"`
	Result  ResponceRowset `xml:"result"`
}


/**
 *	Unused structures
 */

type GetTableByTitleRequest struct {
	XMLName    xml.Name `xml:"http://mfisoft.ru/voip/service/soap getTableByTitle"`
	P_table_hi string `xml:"p_table_hi"`
}

type GetTableByTitleResponce struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap getTableByTitleResponse"`
	Result  string `xml:"result"`
}

type GetColumnLookupRequest struct {
	XMLName     xml.Name `xml:"http://mfisoft.ru/voip/service/soap getColumnLookup"`
	P_table_hi  string `xml:"p_table_hi"`
	P_column_nm string `xml:"p_column_nm"`
}

type GetColumnLookupResponce struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap getColumnLookupResponse"`
	Result  ResponceRowset `xml:"result"` // TODO responce not tested
}


