package rtu

import (
	"github.com/ekrukov/rtu/soap"
)

type RTUClient struct {
	SOAPClient *soap.ServicePortType
}
func NewRTUClient(s, l, p string) *RTUClient{
	client := new(RTUClient)
	clientAuth := &soap.SOAPAuth{Login: l, Password: p}
	client.SOAPClient  = soap.NewServicePortType("https://" + s + "/service/service.php?soap", true, clientAuth)
	return client
}

