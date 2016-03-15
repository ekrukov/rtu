package rtu


import (
	"gamma/rtu/soap"
	"log"
	"errors"
)

type RTUQuery struct {
	client *RTUClient
	Action string
	TableName string
	tableId string
	Filter *soap.Filter
	Rowset *soap.Rowset
	Sort soap.Ordertype
	Limit int
	Offset int
	InsertTemplate *soap.Template
	err string
}

func NewRTUQuery(c *RTUClient) *RTUQuery{
	return &RTUQuery{
		client: c,
		Sort: soap.OrdertypeAsc,
		Limit: 1000,
		Offset: 0,
	}
}

func (q *RTUQuery) Select() *RTUQuery{
	q.Action = "select"
	return q
}

func (q *RTUQuery) Update(table string) *RTUQuery{
	q.Action = "update"
	q.tableId = soap.GetTableIdByName(table)
	return q
}

func (q *RTUQuery) Delete() *RTUQuery{
	q.Action = "delete"
	return q
}

func (q *RTUQuery) Insert(template *soap.Template) *RTUQuery{
	q.Action = "insert"
	q.InsertTemplate = template
	return q
}

func (q *RTUQuery) Set(rowset *soap.Rowset) *RTUQuery{
	if q.Action != "update" {
		q.err = "RTUQuery builder error, set without update"
		log.Fatal(q.err)
	} else {
		q.Rowset = rowset
	}
	return q
}

func (q *RTUQuery) From(table string) *RTUQuery{
	if q.Action == "select" || q.Action == "delete" {
		q.tableId = soap.GetTableIdByName(table)
	} else {
		q.err = "RTUQuery builder error, from without select or delete"
		log.Fatal(q.err)
	}
	return q
}

func (q *RTUQuery) Into(table string) *RTUQuery{
	if q.Action == "insert" {
		q.tableId = soap.GetTableIdByName(table)
	} else {
		q.err = "RTUQuery builder error, into without insert"
		log.Fatal(q.err)
	}
	return q
}

func (q *RTUQuery) Values(rowset *soap.Rowset) *RTUQuery{
	q.Rowset = rowset
	return q
}

func (q *RTUQuery) Where(filter *soap.Filter) *RTUQuery{
	q.Filter = filter
	return q
}

func (q *RTUQuery) OrderBy(sort soap.Ordertype) *RTUQuery{
	q.Sort = sort
	return q
}

func (q *RTUQuery) Run() (res interface{}, err error) {
	switch q.Action {
	case "select":
		res, err = q.client.SOAPClient.SelectRowset(&soap.SelectRowsetRequest{
			P_table_hi: q.tableId,
			Filter: *q.Filter,
			Sort: q.Sort,
			Limit: q.Limit,
		})
		return res, err
	}
	return nil , errors.New("RTUQuery run error action not found")
}

func (q *RTUQuery) Print(){
	log.Printf("%+v", q)
}


