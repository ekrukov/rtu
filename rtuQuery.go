package rtu

import (
	"github.com/ekrukov/rtu/soap"
	"log"
	"errors"
)

type RTUQuery struct {
	client         *RTUClient
	Action         string
	TableName      string
	tableId        string
	Filter         *soap.Filter
	Rowset         *soap.Rowset
	Sort           soap.Ordertype
	Limit          int
	Offset         int
	InsertTemplate *soap.Template
	err            error
}

func NewRTUQuery(c *RTUClient) *RTUQuery {
	return &RTUQuery{
		client: c,
		Sort: soap.OrdertypeAsc,
		Limit: 1000,
		Offset: 0,
	}
}

func (q *RTUQuery) Select() *RTUQuery {
	q.Action = "select"
	return q
}

func (q *RTUQuery) Update(table string) *RTUQuery {
	q.Action = "update"
	q.tableId, q.err = soap.GetTableIdByName(table)
	return q
}

func (q *RTUQuery) Delete() *RTUQuery {
	q.Action = "delete"
	return q
}

func (q *RTUQuery) Insert(template *soap.Template) *RTUQuery {
	q.Action = "insert"
	q.InsertTemplate = template
	return q
}

func (q *RTUQuery) Set(rowset *soap.Rowset) *RTUQuery {
	if q.Action != "update" {
		errorString := "RTUQuery builder error, set without update"
		q.err = errors.New(errorString)
		log.Fatal(errorString)
	} else {
		q.Rowset = rowset
	}
	return q
}

func (q *RTUQuery) From(table string) *RTUQuery {
	if q.Action == "select" || q.Action == "delete" {
		q.tableId, q.err = soap.GetTableIdByName(table)
	} else {
		errorString := "RTUQuery builder error, from without select or delete"
		q.err = errors.New(errorString)
		log.Fatal(errorString)
	}
	return q
}

func (q *RTUQuery) Into(table string) *RTUQuery {
	if q.Action == "insert" {
		q.tableId, q.err = soap.GetTableIdByName(table)
	} else {
		errorString := "RTUQuery builder error, into without insert"
		q.err = errors.New(errorString)
		log.Fatal(errorString)
	}
	return q
}

func (q *RTUQuery) Values(rowset *soap.Rowset) *RTUQuery {
	q.Rowset = rowset
	return q
}

func (q *RTUQuery) Where(filter *soap.Filter) *RTUQuery {
	q.Filter = filter
	return q
}

func (q *RTUQuery) OrderBy(sort soap.Ordertype) *RTUQuery {
	q.Sort = sort
	return q
}

func (q *RTUQuery) Describe(table string) *RTUQuery {
	q.Action = "describe"
	q.tableId, q.err = soap.GetTableIdByName(table)
	return q
}

func (q *RTUQuery) Count(table string, filter *soap.Filter) *RTUQuery {
	q.Action = "count"
	q.tableId, q.err = soap.GetTableIdByName(table)
	q.Filter = filter
	return q
}

func (q *RTUQuery) Run() (res *QueryResponce, err error) {
	res = new(QueryResponce)
	if q.err != nil {
		return nil, q.err
	}
	switch q.Action {
	case "select":
		res.Select, err = q.client.SOAPClient.SelectRowset(&soap.SelectRowsetRequest{
			P_table_hi: q.tableId,
			Filter: *q.Filter,
			Sort: q.Sort,
			Limit: q.Limit,
		})
		return res, err
	case "describe":
		res.Describe, err = q.client.SOAPClient.DescribeColumns(q.tableId)
		return res, err
	case "count":
		count, err := q.client.SOAPClient.CountRowset(&soap.CountRowsetRequest{
			P_table_hi: q.tableId,
			Filter: *q.Filter,
		})
		res.Count = count.Result
		return res, err
	}
	return nil, errors.New("RTUQuery run error action not found")
}

func (q *RTUQuery) Print() {
	log.Printf("%+v", q)
}

type QueryResponce struct {
	Select   *soap.SelectRowsetResponce
	Insert   int
	Delete   int
	Describe *soap.DescribeColumnResponce
	Count    int
}


