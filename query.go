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
	sort           map[string]Ordertype
	limit          int
	offset         int
	insertTemplate map[string]string
	result         *rawResult
	err            error
}

type rawResult struct {
	Describe *DescribeColumnsResponse //rows in result
	Select   *SelectRowsetResponse    //rows in result
	Insert   *InsertRowsetResponse    //int in result
	Delete   *DeleteRowsetResponse    //int in result
	Update   *UpdateRowsetResponse    //int in result
	Count    *CountRowsetResponse     //int in result
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

func (q *RTUQuery) OrderBy(sort map[string]Ordertype) *RTUQuery {
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

func (q *RTUQuery) queryExec() error {
	if q.err != nil {
		return q.err
	}
	q.result = new(rawResult)
	switch q.method {
	case SelectMethod:
		request :=new(SelectRowsetRequest)
		request.requestLimit.P_limit = q.limit
		request.requestOffset.P_offset = q.offset
		if q.tableId == "" {
			return errors.New("need table for select")
		}
		request.requestTable.P_table_hi = q.tableId
		if q.filter != nil {
			filter, err := MapToFilter(q.filter)
			if err != nil {
				return err
			}
			request.requestFilter = *filter
		} else {
			return errors.New("need filter for select")
		}
		if q.sort != nil {
			sort, err := MapToSort(q.sort)
			if err != nil {
				return err
			}
			request.requestSort = *sort
		}

		err := q.client.Call(q.method, &request, &q.result.Select)
		if err != nil {
			return err
		}
	case DescribeMethod:
		request := new(DescribeColumnsRequest)
		request.requestTable.P_table_hi = q.tableId
		err := q.client.Call(q.method, &request, &q.result.Describe)
		if err != nil {
			return err
		}
	case CountMethod:
		filter, err := MapToFilter(q.filter)
		if err != nil {
			return err
		}
		request := new(CountRowsetRequest)
		request.requestTable.P_table_hi = q.tableId
		request.requestFilter = *filter
		err = q.client.Call(q.method, &request, &q.result.Count)
		if err != nil {
			return err
		}
	case InsertMethod:
		rowset, err := MapsToRowset(q.rowset)
		if err != nil {
			return err
		}
		request := new(InsertRowsetRequest)
		request.requestTable.P_table_hi = q.tableId
		request.requestRowset = *rowset
		err = q.client.Call(q.method, &request, &q.result.Insert)
		if err != nil {
			return err
		}
	case UpdateMethod:
		filter, err := MapToFilter(q.filter)
		if err != nil {
			return err
		}
		rowset, err := MapsToRowset(q.rowset)
		if err != nil {
			return err
		}
		request := new(UpdateRowsetRequest)
		request.requestTable.P_table_hi = q.tableId
		request.requestRowset = *rowset
		request.requestFilter = *filter
		err = q.client.Call(q.method, &request, &q.result.Update)
		if err != nil {
			return err
		}
	default:
		return errors.New("RTUQuery run error action not found")
	}
	return nil
}

func (q *RTUQuery) GetRaw() (*rawResult, error) {
	err := q.queryExec()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return q.result, nil
}

func (q *RTUQuery) GetRows() ([]responseRow, error) {
	err := q.queryExec()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	switch q.method {
	case SelectMethod:
		return q.result.Select.Result.Items, nil
	case DescribeMethod:
		return q.result.Describe.Result.Items, nil
	}
	return nil, errors.New("Unsupported method for rows response")
}

func (q *RTUQuery) GetInt() (int, error) {
	err := q.queryExec()
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	switch q.method {
	case InsertMethod:
		return q.result.Insert.Result, nil
	case UpdateMethod:
		return q.result.Update.Result, nil
	case DeleteMethod:
		return q.result.Delete.Result, nil
	case CountMethod:
		return q.result.Count.Result, nil
	}
	return 0, errors.New("Unsupported method for int response")
}

func (q *RTUQuery) GetCDRs() (cs *CDRs, err error) {
	if q.method != SelectMethod{
		err = errors.New("Only select method available for GetCDRs")
		return nil, err
	}
	cs = new(CDRs)
	rows, err := q.GetRows()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for _, it := range rows {
		cdr := new(CDR)
		for _, item := range it.Items {
			cdr.SetField(item.Key, item.Value)
		}
		cs.Items = append(cs.Items, *cdr)
	}
	return cs, nil
}

func (q *RTUQuery) Print() {
	log.Printf("%+v", q)
}





