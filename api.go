package rtu

import (
	"encoding/xml"
)

//
//Request structures
//

type requestColumnName struct {
	P_column_nm string `xml:"p_column_nm"`
}

type requestTable struct {
	P_table_hi string `xml:"p_table_hi"`
}

type requestFilterItem struct {
	XMLName  xml.Name `xml:"http://mfisoft.ru/voip/service/soap p_filter"`
	Type_    string `xml:"type,omitempty"`
	Column   string `xml:"column,omitempty"`
	Operator string `xml:"operator,omitempty"`
	Value    string `xml:"value,omitempty"`
	Childs   *requestFilterChildsArr `xml:"childs,omitempty"`
}

type requestFilter struct {
	Item *requestFilterItem
}

type requestFilterChildsArr struct {
	Items []*requestFilterItem
}

type requestSort struct {
	P_sort struct {
		       XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap p_sort"`
		       Items   []*requestSortItem `xml:"http://mfisoft.ru/voip/service/soap item"`
	       }
}

type requestSortItem struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap sort_item"`
	Column  string `xml:"column,omitempty"`
	Dir     Ordertype `xml:"dir,omitempty"`
}

type requestRowset struct {
	P_rowset struct {
			 XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap p_rowset"`
			 Rows    []*requestRow `xml:"item"`
		 }
}

type requestRow struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap row"`
	Items   []*requestColumn `xml:"http://mfisoft.ru/voip/service/soap item"`
}

type requestColumn struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap column"`
	Name    string `xml:"name,omitempty"`
	Value   string `xml:"value,omitempty"`
}

type requestLimit struct {
	P_limit int `xml:"p_limit"`
}

type requestOffset struct {
	P_offset int `xml:"p_offset"`
}

//
//Response structures
//

type responseString struct {
	Result string `xml:"result"`
}

type responseInt struct {
	Result int `xml:"result"`
}

type responseRowset struct {
	Result struct {
		       XMLName xml.Name `xml:"result"`
		       Items   []responseRow `xml:"item"`
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

type selectRowsetRequest struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap selectRowset"`
	requestTable
	requestFilter
	requestSort
	requestLimit
	requestOffset
}

type selectRowsetResponse struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap selectRowsetResponse"`
	responseRowset
}

/**
 *	Insert Request structures
 */

type insertRowsetRequest struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap insertRowset"`
	requestTable
	requestRowset
}

type insertRowsetResponse struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap insertRowsetResponse"`
	responseInt
}

/**
 *	Update Request structures
 */

type updateRowsetRequest struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap updateRowset"`
	requestTable
	requestRowset
	requestFilter
}

type updateRowsetResponse struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap updateRowsetResponse"`
	responseInt
}
/**
 *	Delete Request structures
 */

type deleteRowsetRequest struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap deleteRowset"`
	requestTable
	requestRowset
	requestFilter
}

type deleteRowsetResponse struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap deleteRowsetResponse"`
	responseInt
}

/**
 *	Count Request structures
 */

type countRowsetRequest struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap countRowset"`
	requestTable
	requestFilter
}

type countRowsetResponse struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap countRowsetResponse"`
	responseInt
}

/**
 *	Describe Request structures
 */

type describeColumnsRequest struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap describeColumns"`
	requestTable
}

type describeColumnsResponse struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap describeColumnsResponse"`
	responseRowset
}


/**
 *	Unused structures
 */

type getTableByTitleRequest struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap getTableByTitle"`
	requestTable
}

type getTableByTitleResponse struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap getTableByTitleResponse"`
	responseString
}

type getColumnLookupRequest struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap getColumnLookup"`
	requestTable
	requestColumnName
}

type getColumnLookupResponse struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap getColumnLookupResponse"`
	responseRowset // TODO response not tested
}


