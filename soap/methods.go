package soap

import "encoding/xml"

type String string
type Integer int
type AnyType interface {}

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

func (service *ServicePortType) SelectRowset(request *SelectRowsetRequest) (*SelectRowsetResponce, error) {
	response := new(SelectRowsetResponce)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePortType) InsertRowset(request *String) (*Integer, error) {
	response := new(Integer)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePortType) UpdateRowset(request *String) (*Integer, error) {
	response := new(Integer)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePortType) DeleteRowset(request *String) (*Integer, error) {
	response := new(Integer)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}



func (service *ServicePortType) CountRowset(request *String) (*Integer, error) {
	response := new(Integer)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePortType) GetTableByTitle(request *String) (*String, error) {
	response := new(String)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type DescribeColumnRequest struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap describeColumns"`
	P_table_hi string `xml:",omitempty"`
}

func (service *ServicePortType) DescribeColumns(request string) (*AnyType, error) {
	response := new(AnyType)
	req := new(DescribeColumnRequest)
	req.P_table_hi = request
	err := service.client.Call("", req, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *ServicePortType) GetColumnLookup(request *String) (*AnyType, error) {
	response := new(AnyType)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type SelectRowsetRequest struct {
	XMLName xml.Name `xml:"http://mfisoft.ru/voip/service/soap selectRowset"`
	P_table_hi string `xml:",omitempty"`
	Filter Filter `xml:",omitempty"`
	Sort Ordertype `xml:",omitempty"`
	Limit int `xml:",omitempty"`
	Offset int `xml:",omitempty"`
}

type SelectRowsetResponce struct {
	XMLName     xml.Name `xml:"http://mfisoft.ru/soap selectRowsetResponse"`
	Result	struct{
			    XMLName xml.Name `xml:"result"`
			    ArrayType string `xml:"http://schemas.xmlsoap.org/soap/encoding/ arrayType,attr"`
			    Type_ string `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`
			    Item struct{
					    XMLName     xml.Name `xml:"item"`
					    Type_       string    `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`
					    Items       []selectRowsetItem `xml:"item"`
				    } `xml:"item"`
		    } `xml:"result"`
}

type selectRowsetItem struct {
	XMLName     xml.Name `xml:"item"`
	Key	    string `xml:"key"`
	Value	    string `xml:"value"`
}

