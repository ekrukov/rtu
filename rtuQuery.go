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
	Filter         map[string]string
	Rowset         *[]map[string]string
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

//TODO template as in param => q.InsertTemplate

func (q *RTUQuery) Insert() *RTUQuery {
	q.Action = "insert"
	return q
}

func (q *RTUQuery) Set(rowset *[]map[string]string) *RTUQuery {
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

func (q *RTUQuery) Values(rowset *[]map[string]string) *RTUQuery {
	q.Rowset = rowset
	return q
}

func (q *RTUQuery) Where(filter map[string]string) *RTUQuery {
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

func (q *RTUQuery) Count(table string, filter map[string]string) *RTUQuery {
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
		filter, err := soap.MapToFilter(q.Filter)
		if err != nil {
			return nil, err
		}
		res.Select, err = q.client.SOAPClient.SelectRowset(&soap.SelectRowsetRequest{
			P_table_hi: q.tableId,
			Filter: *filter,
			Sort: q.Sort,
			Limit: q.Limit,
		})
		if err != nil {
			return nil, err
		}
	case "describe":
		res.Describe, err = q.client.SOAPClient.DescribeColumns(q.tableId)
	case "count":
		filter, err := soap.MapToFilter(q.Filter)
		if err != nil {
			return nil, err
		}
		count, err := q.client.SOAPClient.CountRowset(&soap.CountRowsetRequest{
			P_table_hi: q.tableId,
			Filter: *filter,
		})
		if err != nil {
			return nil, err
		}
		res.Count = count.Result
	case "insert":
		rowset, err := soap.MapsToRowset(q.Rowset)
		if err != nil {
			return nil, err
		}
		insert, err := q.client.SOAPClient.InsertRowset(&soap.InsertRowsetRequest{
			P_table_hi: q.tableId,
			P_rowset: *rowset,
		})
		if err != nil {
			return nil, err
		}
		res.Insert = insert.Result
	case "update":
		filter, err := soap.MapToFilter(q.Filter)
		if err != nil {
			return nil, err
		}
		rowset, err := soap.MapsToRowset(q.Rowset)
		if err != nil {
			return nil, err
		}
		update, err := q.client.SOAPClient.UpdateRowset(&soap.UpdateRowsetRequest{
			P_table_hi: q.tableId,
			P_rowset: *rowset,
			Filter: *filter,
		})
		if err != nil {
			return nil, err
		}
		res.Update = update.Result
	default:
		err = errors.New("RTUQuery run error action not found")
	}
	return res, err
}

func (q *RTUQuery) Print() {
	log.Printf("%+v", q)
}

type QueryResponce struct {
	Describe *soap.DescribeColumnResponce
	Select   *soap.SelectRowsetResponce
	Insert   int
	Delete   int
	Update   int
	Count    int
}


