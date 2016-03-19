package rtu

import ()

type RTUClient struct {
	soapClient *soapClient
}

func NewRTUClient(s, l, p string) *RTUClient {
	client := new(RTUClient)
	clientAuth := &soapAuth{Login: l, Password: p}
	client.soapClient = &soapClient{
		url: "https://" + s + "/service/service.php?soap",
		tls:  true,
		auth: clientAuth,
	}
	return client
}

// Create new query with current client
func (r *RTUClient) Query() *queryBuilder {
	return &queryBuilder{
		client: r.soapClient,
		tableId: "",
		limit: 1000,
		offset: 0,
	}
}


