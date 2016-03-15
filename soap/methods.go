package soap

import "encoding/xml"

// TODO empty request and responce for unused operations
type emptyRequest struct {

}

type emptyResponce struct {

}

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
	Result  struct {
			XMLName   xml.Name `xml:"result"`
			ArrayType string `xml:"http://schemas.xmlsoap.org/soap/encoding/ arrayType,attr"`
			Type_     string `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`
			Rows      []struct {
				XMLName xml.Name `xml:"item"`
				Type_   string    `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`
				Items   []selectRowItem `xml:"item"`
			} `xml:"item"`
		} `xml:"result"`
}

type selectRowItem struct {
	XMLName xml.Name `xml:"item"`
	Key     string `xml:"key"`
	Value   string `xml:"value"`
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
	Result  struct {
			XMLName xml.Name `xml:"result"`
		} `xml:"result"`
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
 *	Update Request method and structures TODO create update structures
 */


func (service *ServicePortType) UpdateRowset(request *emptyRequest) (*emptyResponce, error) {
	response := new(emptyResponce)
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
	Result  struct {
			XMLName xml.Name `xml:"result"`
		} `xml:"result"`
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
	P_rowset   Rowset `xml:",omitempty"`
	Filter     Filter `xml:",omitempty"`
}

type CountRowsetResponce struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap countRowsetResponse"`
	Result  struct {
			XMLName xml.Name `xml:"result"`
		} `xml:"result"`
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
 *	Unused method and structures
 */


func (service *ServicePortType) GetTableByTitle(request *emptyRequest) (*emptyResponce, error) {
	response := new(emptyResponce)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type DescribeColumnRequest struct {
	XMLName    xml.Name `xml:"http://mfisoft.ru/voip/service/soap describeColumns"`
	P_table_hi string `xml:",omitempty"`
}

type DescribeColumnResponce struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/soap describeColumnsResponse"`
	Result  struct {
			XMLName   xml.Name `xml:"result"`
			ArrayType string `xml:"http://schemas.xmlsoap.org/soap/encoding/ arrayType,attr"`
			Type_     string `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`
			Rows      []struct {
				XMLName xml.Name `xml:"item"`
				Type_   string    `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`
				Items   []decribeItem `xml:"item"`
			} `xml:"item"`
		} `xml:"result"`
}

type decribeItem struct {
	XMLName xml.Name `xml:"item"`
	Key     string `xml:"key"`
	Value   string `xml:"value"`
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

func (service *ServicePortType) GetColumnLookup(request *emptyRequest) (*emptyResponce, error) {
	response := new(emptyResponce)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}


