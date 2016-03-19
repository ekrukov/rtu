package rtu

import (
	"encoding/xml"
)

type Template struct {

}

//
//Request structures
//

type requestColumnName struct {
	P_column_nm string `xml:"p_column_nm"`
}

type requestTable struct {
	P_table_hi string `xml:"p_table_hi"`
}

type requestFilter struct {
	P_filter   struct {
			   XMLName  xml.Name `xml:"http://mfisoft.ru/voip/service/soap p_filter"`
			   Type_    string `xml:"type,omitempty"`
			   Column   string `xml:"column,omitempty"`
			   Operator string `xml:"operator,omitempty"`
			   Value    string `xml:"value,omitempty"`
			   //Childs   childFilterArray `xml:"childs,omitempty"`
		   }
}
// TODO only simple filter
/*
type childFilterArray struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap filter_childs_arr"`
	Filters []requestFilter `xml:"http://mfisoft.ru/voip/service/soap item"`
}*/

type requestSort struct {
	P_sort     struct {
			   XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap p_sort"`
			   Items   []requestSortItem `xml:"http://mfisoft.ru/voip/service/soap item"`
		   }
}

type requestSortItem struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap sort_item"`
	Column  string `xml:"column,omitempty"`
	Dir     Ordertype `xml:"dir,omitempty"`
}

type requestRowset struct {
	P_rowset   struct {
			   XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap p_rowset"`
			   Rows    []requestRow `xml:"item"`
		   }
}

type requestRow struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap row"`
	Items   []requestColumn `xml:"http://mfisoft.ru/voip/service/soap item"`
}

type requestColumn struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap column"`
	Name    string `xml:"name,omitempty"`
	Value   string `xml:"value,omitempty"`
}

type requestLimit struct {
	P_limit    int `xml:"p_limit"`
}

type requestOffset struct {
	P_offset   int `xml:"p_offset"`
}

//
//Response structures
//

type responseString struct  {
	Result  string `xml:"result"`
}

type responseInt struct  {
	Result  int `xml:"result"`
}

type responseRowset struct {
	Result  struct{
		XMLName xml.Name `xml:"result"`
		Items []responseRow `xml:"item"`
		}
}

type responseRow struct {
	XMLName xml.Name `xml:"item"`
	Items   []responseColumn `xml:"item"`
}

type responseColumn struct {
	XMLName xml.Name `xml:"item"`
	Key     string `xml:"key"`
	Value   string `xml:"value"`
}

//
// Method structures
//

type SelectRowsetRequest struct {
	XMLName    xml.Name `xml:"http://mfisoft.ru/voip/service/soap selectRowset"`
	requestTable
	requestFilter
	requestSort
	requestLimit
	requestOffset
}

type SelectRowsetResponse struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap selectRowsetResponse"`
	responseRowset
}

/**
 *	Insert Request structures
 */

type InsertRowsetRequest struct {
	XMLName    xml.Name `xml:"http://mfisoft.ru/voip/service/soap insertRowset"`
	requestTable
	requestRowset
}

type InsertRowsetResponse struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap insertRowsetResponse"`
	responseInt
}

/**
 *	Update Request structures
 */

type UpdateRowsetRequest struct {
	XMLName    xml.Name `xml:"http://mfisoft.ru/voip/service/soap updateRowset"`
	requestTable
	requestRowset
	requestFilter
}

type UpdateRowsetResponse struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap updateRowsetResponse"`
	responseInt
}
/**
 *	Delete Request structures
 */

type DeleteRowsetRequest struct {
	XMLName    xml.Name `xml:"http://mfisoft.ru/voip/service/soap deleteRowset"`
	requestTable
	requestRowset
	requestFilter
}

type DeleteRowsetResponse struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap deleteRowsetResponse"`
	responseInt
}

/**
 *	Count Request structures
 */

type CountRowsetRequest struct {
	XMLName    xml.Name `xml:"http://mfisoft.ru/voip/service/soap countRowset"`
	requestTable
	requestFilter
}

type CountRowsetResponse struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap countRowsetResponse"`
	responseInt
}

/**
 *	Describe Request structures
 */

type DescribeColumnsRequest struct {
	XMLName    xml.Name `xml:"http://mfisoft.ru/voip/service/soap describeColumns"`
	requestTable
}

type DescribeColumnsResponse struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap describeColumnsResponse"`
	responseRowset
}


/**
 *	Unused structures
 */

type GetTableByTitleRequest struct {
	XMLName    xml.Name `xml:"http://mfisoft.ru/voip/service/soap getTableByTitle"`
	requestTable
}

type GetTableByTitleResponse struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap getTableByTitleResponse"`
	responseString
}

type GetColumnLookupRequest struct {
	XMLName     xml.Name `xml:"http://mfisoft.ru/voip/service/soap getColumnLookup"`
	requestTable
	requestColumnName
}

type GetColumnLookupResponse struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap getColumnLookupResponse"`
	responseRowset // TODO response not tested
}


