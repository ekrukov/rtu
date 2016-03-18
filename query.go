package rtu

import (
	"log"
	"errors"
)

type RTUQuery struct {
	client         *SOAPClient
	method         Methodtype
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
	q.method = SelectMethod
	return q
}

func (q *RTUQuery) Update(table string) *RTUQuery {
	q.method = UpdateMethod
	q.tableId, q.err = GetTableIdByName(table)
	return q
}

func (q *RTUQuery) Delete() *RTUQuery {
	q.method = DeleteMethod
	return q
}

//TODO template as in param => q.InsertTemplate

func (q *RTUQuery) Insert() *RTUQuery {
	q.method = InsertMethod
	return q
}

func (q *RTUQuery) Set(rowset []map[string]string) *RTUQuery {
	if q.method != UpdateMethod {
		errorString := "RTUQuery builder error, set without update"
		q.err = errors.New(errorString)
		log.Fatal(errorString)
	} else {
		q.rowset = rowset
	}
	return q
}

func (q *RTUQuery) From(table string) *RTUQuery {
	if q.method == SelectMethod || q.method == DeleteMethod {
		q.tableId, q.err = GetTableIdByName(table)
	} else {
		errorString := "RTUQuery builder error, from without select or delete"
		q.err = errors.New(errorString)
		log.Fatal(errorString)
	}
	return q
}

func (q *RTUQuery) Into(table string) *RTUQuery {
	if q.method == InsertMethod {
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
	q.method = DescribeMethod
	q.tableId, q.err = GetTableIdByName(table)
	return q
}

func (q *RTUQuery) Count(table string, filter map[string]string) *RTUQuery {
	q.method = CountMethod
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
	case SelectMethod:
		request := SelectRowsetRequest{
			P_limit: q.limit,
			P_offset: q.offset,
		}
		if q.tableId == "" {
			return nil, errors.New("need table for select")
		}
		request.P_table_hi = q.tableId
		if q.filter != nil {
			filter, err := MapToFilter(q.filter)
			if err != nil {
				return nil, err
			}
			request.P_filter = *filter
		} else {
			return nil, errors.New("need filter for select")
		}
		if q.sort != nil {
			sort, err := MapToSort(q.sort)
			if err != nil {
				return nil, err
			}
			request.P_sort = *sort
		}

		err = q.client.Call(q.method, &request, &res.Select)
		if err != nil {
			return nil, err
		}
	case DescribeMethod:
		err = q.client.Call(q.method, &DescribeColumnsRequest{
			P_table_hi: q.tableId,
		}, &res.Describe)
		if err != nil {
			return nil, err
		}
	case CountMethod:
		filter, err := MapToFilter(q.filter)
		if err != nil {
			return nil, err
		}
		count := new(CountRowsetResponce)
		err = q.client.Call(q.method, &CountRowsetRequest{
			P_table_hi: q.tableId,
			Filter: *filter,
		}, &count)
		if err != nil {
			return nil, err
		}
		res.Count = count.Result
	case InsertMethod:
		rowset, err := MapsToRowset(q.rowset)
		if err != nil {
			return nil, err
		}
		insert := new(InsertRowsetResponce)
		err = q.client.Call(q.method, &InsertRowsetRequest{
			P_table_hi: q.tableId,
			P_rowset: *rowset,
		}, &insert)
		if err != nil {
			return nil, err
		}
		res.Insert = insert.Result
	case UpdateMethod:
		filter, err := MapToFilter(q.filter)
		if err != nil {
			return nil, err
		}
		rowset, err := MapsToRowset(q.rowset)
		if err != nil {
			return nil, err
		}
		update := new(UpdateRowsetResponce)
		err = q.client.Call(q.method, &UpdateRowsetRequest{
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
	Describe *DescribeColumnsResponce
	Select   *SelectRowsetResponce
	Insert   int
	Delete   int
	Update   int
	Count    int
}




