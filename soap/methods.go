package soap

import "encoding/xml"

type SOAPService struct {
	client *SOAPClient
}

func NewSOAPService(url string, tls bool, auth *SOAPAuth) *SOAPService {
	if url == "" {
		url = ""
	}
	client := NewSOAPClient(url, tls, auth)

	return &SOAPService{
		client: client,
	}
}

/**
 *	Select Request method and structures
 */

type SelectRowsetRequest struct {
	XMLName    xml.Name `xml:"http://mfisoft.ru/voip/service/soap selectRowset"`
	P_table_hi string `xml:",omitempty"`
	Filter     Filter `xml:",omitempty"`
	Sort       Ordertype `xml:",omitempty"`
	Limit      int `xml:",omitempty"`
	Offset     int `xml:",omitempty"`
}

type SelectRowsetResponce struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap selectRowsetResponse"`
	Result  ResponceRowset `xml:"result"`
}

func (service *SOAPService) SelectRowset(request *SelectRowsetRequest) (*SelectRowsetResponce, error) {
	response := new(SelectRowsetResponce)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/**
 *	Insert Request method and structures
 */

type InsertRowsetRequest struct {
	XMLName    xml.Name `xml:"http://mfisoft.ru/voip/service/soap insertRowset"`
	P_table_hi string `xml:",omitempty"`
	P_rowset   Rowset `xml:",omitempty"`
}

type InsertRowsetResponce struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap insertRowsetResponse"`
	Result  int `xml:"result"`
}

func (service *SOAPService) InsertRowset(request *InsertRowsetRequest) (*InsertRowsetResponce, error) {
	response := new(InsertRowsetResponce)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/**
 *	Update Request method and structures
 */

type UpdateRowsetRequest struct {
	XMLName    xml.Name `xml:"http://mfisoft.ru/voip/service/soap updateRowset"`
	P_table_hi string `xml:",omitempty"`
	P_rowset   Rowset `xml:",omitempty"`
	Filter     Filter `xml:",omitempty"`
}

type UpdateRowsetResponce struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap updateRowsetResponse"`
	Result  int `xml:"result"`
}

func (service *SOAPService) UpdateRowset(request *UpdateRowsetRequest) (*UpdateRowsetResponce, error) {
	response := new(UpdateRowsetResponce)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/**
 *	Delete Request method and structures
 */

type DeleteRowsetRequest struct {
	XMLName    xml.Name `xml:"http://mfisoft.ru/voip/service/soap deleteRowset"`
	P_table_hi string `xml:",omitempty"`
	P_rowset   Rowset `xml:",omitempty"`
	Filter     Filter `xml:",omitempty"`
}

type DeleteRowsetResponce struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap deleteRowsetResponse"`
	Result  int `xml:"result"`
}

func (service *SOAPService) DeleteRowset(request *DeleteRowsetRequest) (*DeleteRowsetResponce, error) {
	response := new(DeleteRowsetResponce)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/**
 *	Count Request method and structures
 */

type CountRowsetRequest struct {
	XMLName    xml.Name `xml:"http://mfisoft.ru/voip/service/soap countRowset"`
	P_table_hi string `xml:",omitempty"`
	Filter     Filter `xml:",omitempty"`
}

type CountRowsetResponce struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap countRowsetResponse"`
	Result  int `xml:"result"`
}

func (service *SOAPService) CountRowset(request *CountRowsetRequest) (*CountRowsetResponce, error) {
	response := new(CountRowsetResponce)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/**
 *	Describe Request method and structures
 */

type DescribeColumnRequest struct {
	XMLName    xml.Name `xml:"http://mfisoft.ru/voip/service/soap describeColumns"`
	P_table_hi string `xml:",omitempty"`
}

type DescribeColumnResponce struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap describeColumnsResponse"`
	Result  ResponceRowset `xml:"result"`
}

func (service *SOAPService) DescribeColumns(request *DescribeColumnRequest) (*DescribeColumnResponce, error) {
	response := new(DescribeColumnResponce)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}


/**
 *	Unused method and structures
 */

type GetTableByTitleRequest struct {
	XMLName    xml.Name `xml:"http://mfisoft.ru/voip/service/soap getTableByTitle"`
	P_table_hi string `xml:",omitempty"`
}

type GetTableByTitleResponce struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap getTableByTitleResponse"`
	Result  string `xml:"result"`
}

func (service *SOAPService) GetTableByTitle(request *GetTableByTitleRequest) (*GetTableByTitleResponce, error) {
	response := new(GetTableByTitleResponce)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type GetColumnLookupRequest struct {
	XMLName     xml.Name `xml:"http://mfisoft.ru/voip/service/soap getColumnLookup"`
	P_table_hi  string `xml:",omitempty"`
	P_column_nm string `xml:",omitempty"`
}

type GetColumnLookupResponce struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap getColumnLookupResponse"`
	Result  ResponceRowset `xml:"result"` // TODO responce not tested
}

func (service *SOAPService) GetColumnLookup(request *GetColumnLookupRequest) (*GetColumnLookupResponce, error) {
	response := new(GetColumnLookupResponce)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
