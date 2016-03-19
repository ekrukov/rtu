package rtu

import (
	"log"
	"errors"
)

type queryBuilder struct {
	client         *soapClient
	method         methodType
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
	Describe *describeColumnsResponse //rows in result
	Select   *selectRowsetResponse    //rows in result
	Insert   *insertRowsetResponse    //int in result
	Delete   *deleteRowsetResponse    //int in result
	Update   *updateRowsetResponse    //int in result
	Count    *countRowsetResponse     //int in result
}

func (q *queryBuilder) Select() *queryBuilder {
	q.method = selectMethod
	return q
}

func (q *queryBuilder) Update(table string) *queryBuilder {
	q.method = updateMethod
	q.tableId, q.err = GetTableIdByName(table)
	return q
}

func (q *queryBuilder) Delete() *queryBuilder {
	q.method = deleteMethod
	return q
}

//TODO template as in param => q.InsertTemplate

func (q *queryBuilder) Insert() *queryBuilder {
	q.method = insertMethod
	return q
}

func (q *queryBuilder) Set(rowset []map[string]string) *queryBuilder {
	if q.method != updateMethod {
		errorString := "queryBuilder builder error, set without update"
		q.err = errors.New(errorString)
		log.Fatal(errorString)
	} else {
		q.rowset = rowset
	}
	return q
}

func (q *queryBuilder) From(table string) *queryBuilder {
	if q.method == selectMethod || q.method == deleteMethod {
		q.tableId, q.err = GetTableIdByName(table)
	} else {
		errorString := "queryBuilder builder error, from without select or delete"
		q.err = errors.New(errorString)
		log.Fatal(errorString)
	}
	return q
}

func (q *queryBuilder) Into(table string) *queryBuilder {
	if q.method == insertMethod {
		q.tableId, q.err = GetTableIdByName(table)
	} else {
		errorString := "queryBuilder builder error, into without insert"
		q.err = errors.New(errorString)
		log.Fatal(errorString)
	}
	return q
}

func (q *queryBuilder) Values(rowset []map[string]string) *queryBuilder {
	q.rowset = rowset
	return q
}

func (q *queryBuilder) Where(filter map[string]string) *queryBuilder {
	q.filter = filter
	return q
}

func (q *queryBuilder) OrderBy(sort map[string]Ordertype) *queryBuilder {
	q.sort = sort
	return q
}

func (q *queryBuilder) Limit(limit int) *queryBuilder {
	q.limit = limit
	return q
}

func (q *queryBuilder) Offset(offset int) *queryBuilder {
	q.offset = offset
	return q
}

func (q *queryBuilder) Describe(table string) *queryBuilder {
	q.method = describeMethod
	q.tableId, q.err = GetTableIdByName(table)
	return q
}

func (q *queryBuilder) Count(table string, filter map[string]string) *queryBuilder {
	q.method = countMethod
	q.tableId, q.err = GetTableIdByName(table)
	q.filter = filter
	return q
}

func (q *queryBuilder) queryExec() error {
	if q.err != nil {
		return q.err
	}
	q.result = new(rawResult)
	switch q.method {
	case selectMethod:
		request :=new(selectRowsetRequest)
		request.requestLimit.P_limit = q.limit
		request.requestOffset.P_offset = q.offset
		if q.tableId == "" {
			return errors.New("need table for select")
		}
		request.requestTable.P_table_hi = q.tableId
		if q.filter != nil {
			filter, err := mapToFilter(q.filter)
			if err != nil {
				return err
			}
			request.requestFilter = *filter
		} else {
			return errors.New("need filter for select")
		}
		if q.sort != nil {
			sort, err := mapToSort(q.sort)
			if err != nil {
				return err
			}
			request.requestSort = *sort
		}

		err := q.client.Call(q.method, &request, &q.result.Select)
		if err != nil {
			return err
		}
	case describeMethod:
		request := new(describeColumnsRequest)
		request.requestTable.P_table_hi = q.tableId
		err := q.client.Call(q.method, &request, &q.result.Describe)
		if err != nil {
			return err
		}
	case countMethod:
		filter, err := mapToFilter(q.filter)
		if err != nil {
			return err
		}
		request := new(countRowsetRequest)
		request.requestTable.P_table_hi = q.tableId
		request.requestFilter = *filter
		err = q.client.Call(q.method, &request, &q.result.Count)
		if err != nil {
			return err
		}
	case insertMethod:
		rowset, err := mapsToRowset(q.rowset)
		if err != nil {
			return err
		}
		request := new(insertRowsetRequest)
		request.requestTable.P_table_hi = q.tableId
		request.requestRowset = *rowset
		err = q.client.Call(q.method, &request, &q.result.Insert)
		if err != nil {
			return err
		}
	case updateMethod:
		filter, err := mapToFilter(q.filter)
		if err != nil {
			return err
		}
		rowset, err := mapsToRowset(q.rowset)
		if err != nil {
			return err
		}
		request := new(updateRowsetRequest)
		request.requestTable.P_table_hi = q.tableId
		request.requestRowset = *rowset
		request.requestFilter = *filter
		err = q.client.Call(q.method, &request, &q.result.Update)
		if err != nil {
			return err
		}
	default:
		return errors.New("queryBuilder run error action not found")
	}
	return nil
}

func (q *queryBuilder) GetRaw() (*rawResult, error) {
	err := q.queryExec()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return q.result, nil
}

func (q *queryBuilder) GetRows() ([]responseRow, error) {
	err := q.queryExec()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	switch q.method {
	case selectMethod:
		return q.result.Select.Result.Items, nil
	case describeMethod:
		return q.result.Describe.Result.Items, nil
	}
	return nil, errors.New("Unsupported method for rows response")
}

func (q *queryBuilder) GetInt() (int, error) {
	err := q.queryExec()
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	switch q.method {
	case insertMethod:
		return q.result.Insert.Result, nil
	case updateMethod:
		return q.result.Update.Result, nil
	case deleteMethod:
		return q.result.Delete.Result, nil
	case countMethod:
		return q.result.Count.Result, nil
	}
	return 0, errors.New("Unsupported method for int response")
}

func (q *queryBuilder) GetCDRs() (cs *CDRs, err error) {
	if q.method != selectMethod{
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

func (q *queryBuilder) Print() {
	log.Printf("%+v", q)
}





