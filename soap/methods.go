package soap

import "encoding/xml"

type ServicePortType struct {
	client *SOAPClient
}

func NewServicePortType(url string, tls bool, auth *SOAPAuth) *ServicePortType {
	if url == "" {
		url = ""
	}
	client := NewSOAPClient(url, tls, auth)

	return &ServicePortType{
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
	Result  ResponceRowResult `xml:"result"`
}

func (service *ServicePortType) SelectRowset(request *SelectRowsetRequest) (*SelectRowsetResponce, error) {
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

func (service *ServicePortType) InsertRowset(request *InsertRowsetRequest) (*InsertRowsetResponce, error) {
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

func (service *ServicePortType) UpdateRowset(request *UpdateRowsetRequest) (*UpdateRowsetResponce, error) {
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

func (service *ServicePortType) DeleteRowset(request *DeleteRowsetRequest) (*DeleteRowsetResponce, error) {
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

func (service *ServicePortType) CountRowset(request *CountRowsetRequest) (*CountRowsetResponce, error) {
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
	Result  ResponceRowResult `xml:"result"`
}

func (service *ServicePortType) DescribeColumns(request string) (*DescribeColumnResponce, error) {
	response := new(DescribeColumnResponce)
	req := new(DescribeColumnRequest)
	req.P_table_hi = request
	err := service.client.Call("", req, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}


/**
 *	Unused method and structures
 *


func (service *ServicePortType) GetTableByTitle(request *emptyRequest) (*emptyResponce, error) {
	response := new(emptyResponce)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePortType) GetColumnLookup(request *emptyRequest) (*emptyResponce, error) {
	response := new(emptyResponce)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}


*/