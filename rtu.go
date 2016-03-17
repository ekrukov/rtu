package rtu

import (
	"github.com/ekrukov/rtu/soap"
)

type RTUClient struct {
	SOAPClient *soap.SOAPClient
}

func NewRTUClient(s, l, p string) *RTUClient {
	client := new(RTUClient)
	clientAuth := &soap.SOAPAuth{Login: l, Password: p}
	client.SOAPClient = soap.NewSOAPClient("https://" + s + "/service/service.php?soap", true, clientAuth)
	return client
}

func (r *RTUClient) Query() *RTUQuery {
	return &RTUQuery{
		client: r.SOAPClient,
		tableId: "",
		limit: 1000,
		offset: 0,
	}
}
