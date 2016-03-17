package rtu

import (
	"github.com/ekrukov/rtu/soap"
	"log"
	"errors"
)

type RTUQuery struct {
	client         *soap.SOAPClient
	method         soap.Methodtype
	tableName      string
	tableId        string
	filter         map[string]string
	rowset         []map[string]string
	sort           map[string]string
	limit          int
	offset         int
	insertTemplate map[string]string
	err            error
}

func (q *RTUQuery) Select() *RTUQuery {
	q.method = soap.SelectMethod
	return q
}

func (q *RTUQuery) Update(table string) *RTUQuery {
	q.method = soap.UpdateMethod
	q.tableId, q.err = GetTableIdByName(table)
	return q
}

func (q *RTUQuery) Delete() *RTUQuery {
	q.method = soap.DeleteMethod
	return q
}

//TODO template as in param => q.InsertTemplate

func (q *RTUQuery) Insert() *RTUQuery {
	q.method = soap.InsertMethod
	return q
}

func (q *RTUQuery) Set(rowset []map[string]string) *RTUQuery {
	if q.method != soap.UpdateMethod {
		errorString := "RTUQuery builder error, set without update"
		q.err = errors.New(errorString)
		log.Fatal(errorString)
	} else {
		q.rowset = rowset
	}
	return q
}

func (q *RTUQuery) From(table string) *RTUQuery {
	if q.method == soap.SelectMethod || q.method == soap.DeleteMethod {
		q.tableId, q.err = GetTableIdByName(table)
	} else {
		errorString := "RTUQuery builder error, from without select or delete"
		q.err = errors.New(errorString)
		log.Fatal(errorString)
	}
	return q
}

func (q *RTUQuery) Into(table string) *RTUQuery {
	if q.method == soap.InsertMethod {
		q.tableId, q.err = GetTableIdByName(table)
	} else {
		errorString := "RTUQuery builder error, into without insert"
		q.err = errors.New(errorString)
		log.Fatal(errorString)
	}
	return q
}

func (q *RTUQuery) Values(rowset []map[string]string) *RTUQuery {
	q.rowset = rowset
	return q
}

func (q *RTUQuery) Where(filter map[string]string) *RTUQuery {
	q.filter = filter
	return q
}

func (q *RTUQuery) OrderBy(sort map[string]string) *RTUQuery {
	q.sort = sort
	return q
}

func (q *RTUQuery) Limit(limit int) *RTUQuery {
	q.limit = limit
	return q
}

func (q *RTUQuery) Offset(offset int) *RTUQuery {
	q.offset = offset
	return q
}

func (q *RTUQuery) Describe(table string) *RTUQuery {
	q.method = soap.DescribeMethod
	q.tableId, q.err = GetTableIdByName(table)
	return q
}

func (q *RTUQuery) Count(table string, filter map[string]string) *RTUQuery {
	q.method = soap.CountMethod
	q.tableId, q.err = GetTableIdByName(table)
	q.filter = filter
	return q
}

func (q *RTUQuery) Run() (res *QueryResponce, err error) {
	res = new(QueryResponce)
	if q.err != nil {
		return nil, q.err
	}
	switch q.method {
	case soap.SelectMethod:
		request := soap.SelectRowsetRequest{
			P_limit: q.limit,
			P_offset: q.offset,
		}
		if q.tableId == "" {
			return nil, errors.New("need table for select")
		}
		request.P_table_hi = q.tableId
		if q.filter != nil {
			filter, err := soap.MapToFilter(q.filter)
			if err != nil {
				return nil, err
			}
			request.P_filter = *filter
		} else {
			return nil, errors.New("need filter for select")
		}
		if q.sort != nil {
			sort, err := soap.MapToSort(q.sort)
			if err != nil {
				return nil, err
			}
			request.P_sort = *sort
		}

		err = q.client.Call(q.method, &request, &res.Select)
		if err != nil {
			return nil, err
		}
	case soap.DescribeMethod:
		err = q.client.Call(q.method, &soap.DescribeColumnsRequest{
			P_table_hi: q.tableId,
		}, &res.Describe)
		if err != nil {
			return nil, err
		}
	case soap.CountMethod:
		filter, err := soap.MapToFilter(q.filter)
		if err != nil {
			return nil, err
		}
		count := new(soap.CountRowsetResponce)
		err = q.client.Call(q.method, &soap.CountRowsetRequest{
			P_table_hi: q.tableId,
			Filter: *filter,
		}, &count)
		if err != nil {
			return nil, err
		}
		res.Count = count.Result
	case soap.InsertMethod:
		rowset, err := soap.MapsToRowset(q.rowset)
		if err != nil {
			return nil, err
		}
		insert := new(soap.InsertRowsetResponce)
		err = q.client.Call(q.method, &soap.InsertRowsetRequest{
			P_table_hi: q.tableId,
			P_rowset: *rowset,
		}, &insert)
		if err != nil {
			return nil, err
		}
		res.Insert = insert.Result
	case soap.UpdateMethod:
		filter, err := soap.MapToFilter(q.filter)
		if err != nil {
			return nil, err
		}
		rowset, err := soap.MapsToRowset(q.rowset)
		if err != nil {
			return nil, err
		}
		update := new(soap.UpdateRowsetResponce)
		err = q.client.Call(q.method, &soap.UpdateRowsetRequest{
			P_table_hi: q.tableId,
			P_rowset: *rowset,
			Filter: *filter,
		}, &update)
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
	Describe *soap.DescribeColumnsResponce
	Select   *soap.SelectRowsetResponce
	Insert   int
	Delete   int
	Update   int
	Count    int
}




