package rtu

type RTUClient struct {
	soapClient *soapClient
}

func NewRTUClient(s, l, p string) *RTUClient {
	client := new(RTUClient)
	clientAuth := &soapAuth{Login: l, Password: p}
	client.soapClient = &soapClient{
		url: urlMake(s),
		tls:  true,
		auth: clientAuth,
	}
	return client
}

// Create new query with current client
func (r *RTUClient) Query() *queryBuilder {
	defaultRequest := requestData{
		Table: &requestTable{},
		Limit: &requestLimit{
			P_limit: queryDefaultLimit,
		},
		Offset: &requestOffset{
			P_offset: 0,
		},
	}
	return &queryBuilder{
		Client: r.soapClient,
		Request: &defaultRequest,
	}
}


