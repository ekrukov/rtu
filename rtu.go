package rtu

import (

)

type RTUClient struct {
	SOAPClient *SOAPClient
}

func NewRTUClient(s, l, p string) *RTUClient {
	client := new(RTUClient)
	clientAuth := &SOAPAuth{Login: l, Password: p}
	client.SOAPClient = NewSOAPClient("https://" + s + "/service/service.php?soap", true, clientAuth)
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
