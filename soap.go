package rtu

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"net"
)
type Template struct {

}

var timeout = time.Duration(30 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

type SOAPClient struct {
	url  string
	tls  bool
	auth *SOAPAuth
}

func NewSOAPClient(url string, tls bool, auth *SOAPAuth) *SOAPClient {
	return &SOAPClient{
		url:  url,
		tls:  tls,
		auth: auth,
	}
}

func (s *SOAPClient) Call(soapMethod Methodtype, request, response interface{}) error {
	envelope := SOAPEnvelope{}

	envelope.Header.Auth.Login = s.auth.Login
	envelope.Header.Auth.Password = s.auth.Password
	envelope.Body.Content = request

	buffer := new(bytes.Buffer)

	encoder := xml.NewEncoder(buffer)
	encoder.Indent("  ", "    ")

	if err := encoder.Encode(envelope); err != nil {
		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}

	//log.Println(buffer.String())

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}
	if s.auth != nil {
		req.SetBasicAuth(s.auth.Login, s.auth.Password)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	if Methodtype(soapMethod) != "" {
		req.Header.Add("SOAPAction", string(soapMethod))
	}

	req.Header.Set("User-Agent", "gowsdl/0.1")
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: s.tls,
		},
		Dial: dialTimeout,
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if len(rawbody) == 0 {
		log.Println("empty response")
		return nil
	}

	//log.Println(string(rawbody))
	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}
	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}

	return nil
}


type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  SOAPHeader
	Body    SOAPBody
}

type SOAPHeader struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`
	Auth    SOAPAuthHeader
}

type SOAPAuthHeader struct {
	XMLName  xml.Name `xml:"http://mfisoft.ru/auth Auth"`
	Login    string `xml:",omitempty"`
	Password string `xml:",omitempty"`
}

type SOAPAuth struct {
	Login    string
	Password string
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`
	Code    string `xml:"faultcode,omitempty"`
	String  string `xml:"faultstring,omitempty"`
	Actor   string `xml:"faultactor,omitempty"`
	Detail  string `xml:"detail,omitempty"`
}

func (f *SOAPFault) Error() string {
	return f.String
}

type SOAPBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token xml.Token
		err error
		consumed bool
	)

	Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
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


