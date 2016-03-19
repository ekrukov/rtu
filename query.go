package rtu

import (
	"log"
	"errors"
)

type queryBuilder struct {
	Client  *soapClient
	Request *requestData
	Result  *rawResult
	err     error
}

type rawResult struct {
	Describe *describeColumnsResponse //rows in result
	Select   *selectRowsetResponse    //rows in result
	Insert   *insertRowsetResponse    //int in result
	Delete   *deleteRowsetResponse    //int in result
	Update   *updateRowsetResponse    //int in result
	Count    *countRowsetResponse     //int in result
}

type requestData struct {
	Method methodType
	Table  *requestTable
	Filter *requestFilter
	Rowset *requestRowset
	Sort   *requestSort
	Limit  *requestLimit
	Offset *requestOffset
}

func (q *queryBuilder) Select() *queryBuilder {
	q.Request.Method = selectMethod
	return q
}

func (q *queryBuilder) Update(table string) *queryBuilder {
	q.Request.Method = updateMethod
	q.Request.Table.P_table_hi, q.err = GetTableIdByName(table)
	return q
}

func (q *queryBuilder) Delete() *queryBuilder {
	q.Request.Method = deleteMethod
	return q
}

//TODO template as in param => q.InsertTemplate

func (q *queryBuilder) Insert() *queryBuilder {
	q.Request.Method = insertMethod
	return q
}

func (q *queryBuilder) Set(rowset []map[string]string) *queryBuilder {
	if q.Request.Method != updateMethod {
		errorString := "queryBuilder builder error, set without update"
		q.err = errors.New(errorString)
		log.Fatal(errorString)
	} else {
		q.Request.Rowset, q.err = mapsToRowset(rowset)
	}
	return q
}

func (q *queryBuilder) From(table string) *queryBuilder {
	if q.Request.Method == selectMethod || q.Request.Method == deleteMethod {
		q.Request.Table.P_table_hi, q.err = GetTableIdByName(table)
	} else {
		errorString := "query builder error: from without select or delete"
		q.err = errors.New(errorString)
		log.Fatal(errorString)
	}
	return q
}

func (q *queryBuilder) Into(table string) *queryBuilder {
	if q.Request.Method == insertMethod {
		q.Request.Table.P_table_hi, q.err = GetTableIdByName(table)
	} else {
		errorString := "queryBuilder builder error, into without insert"
		q.err = errors.New(errorString)
		log.Fatal(errorString)
	}
	return q
}

func (q *queryBuilder) Values(rowset []map[string]string) *queryBuilder {
	q.Request.Rowset, q.err = mapsToRowset(rowset)
	return q
}

func (q *queryBuilder) Where(filter map[string]string) *queryBuilder {
	q.Request.Filter, q.err = mapToFilter(filter)
	return q
}

func (q *queryBuilder) OrderBy(sort map[string]Ordertype) *queryBuilder {
	q.Request.Sort, q.err = mapToSort(sort)
	return q
}

func (q *queryBuilder) Limit(limit int) *queryBuilder {
	q.Request.Limit.P_limit = limit
	return q
}

func (q *queryBuilder) Offset(offset int) *queryBuilder {
	q.Request.Offset.P_offset = offset
	return q
}

func (q *queryBuilder) Describe(table string) *queryBuilder {
	q.Request.Method = describeMethod
	q.Request.Table.P_table_hi, q.err = GetTableIdByName(table)
	return q
}

func (q *queryBuilder) Count(table string, filter map[string]string) *queryBuilder {
	q.Request.Method = countMethod
	q.Request.Table.P_table_hi, q.err = GetTableIdByName(table)
	q.Request.Filter, q.err = mapToFilter(filter)
	return q
}

func (q *queryBuilder) queryExec() error {
	if q.err != nil {
		return q.err
	}
	q.Result = new(rawResult)
	switch q.Request.Method {
	case selectMethod:
		request := new(selectRowsetRequest)
		request.requestLimit = *q.Request.Limit
		request.requestOffset = *q.Request.Offset
		if q.Request.Table != nil {
			request.requestTable = *q.Request.Table
		} else {
			return errors.New("need table name for request")
		}
		if q.Request.Filter != nil {
			request.requestFilter = *q.Request.Filter
		} else {
			return errors.New("need filter for select")
		}
		if q.Request.Sort != nil {
			request.requestSort = *q.Request.Sort
		}

		err := q.Client.Call(q.Request.Method, &request, &q.Result.Select)
		if err != nil {
			return err
		}
	case describeMethod:
		request := new(describeColumnsRequest)
		if q.Request.Table != nil {
			request.requestTable = *q.Request.Table
		} else {
			return errors.New("need table name for request")
		}
		err := q.Client.Call(q.Request.Method, &request, &q.Result.Describe)
		if err != nil {
			return err
		}
	case countMethod:
		request := new(countRowsetRequest)
		if q.Request.Filter != nil {
			request.requestFilter = *q.Request.Filter
		} else {
			return errors.New("need filter for count")
		}
		if q.Request.Table != nil {
			request.requestTable = *q.Request.Table
		} else {
			return errors.New("need table name for request")
		}
		err := q.Client.Call(q.Request.Method, &request, &q.Result.Count)
		if err != nil {
			return err
		}
	case insertMethod:
		request := new(insertRowsetRequest)
		if q.Request.Rowset != nil {
			request.requestRowset = *q.Request.Rowset
		} else {
			return errors.New("need rowset for insert")
		}
		if q.Request.Table != nil {
			request.requestTable = *q.Request.Table
		} else {
			return errors.New("need table name for request")
		}
		err := q.Client.Call(q.Request.Method, &request, &q.Result.Insert)
		if err != nil {
			return err
		}
	case updateMethod:
		request := new(updateRowsetRequest)
		if q.Request.Rowset != nil {
			request.requestRowset = *q.Request.Rowset
		} else {
			return errors.New("need rowset for update")
		}
		if q.Request.Table != nil {
			request.requestTable = *q.Request.Table
		} else {
			return errors.New("need table name for request")
		}
		if q.Request.Filter != nil {
			request.requestFilter = *q.Request.Filter
		} else {
			return errors.New("need filter for count")
		}
		err := q.Client.Call(q.Request.Method, &request, &q.Result.Update)
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
	return q.Result, nil
}

func (q *queryBuilder) GetRows() ([]responseRow, error) {
	err := q.queryExec()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	switch q.Request.Method {
	case selectMethod:
		return q.Result.Select.Result.Items, nil
	case describeMethod:
		return q.Result.Describe.Result.Items, nil
	}
	return nil, errors.New("Unsupported method for rows response")
}

func (q *queryBuilder) GetInt() (int, error) {
	err := q.queryExec()
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	switch q.Request.Method {
	case insertMethod:
		return q.Result.Insert.Result, nil
	case updateMethod:
		return q.Result.Update.Result, nil
	case deleteMethod:
		return q.Result.Delete.Result, nil
	case countMethod:
		return q.Result.Count.Result, nil
	}
	return 0, errors.New("Unsupported method for int response")
}

func (q *queryBuilder) GetCDRs() (cs *CDRs, err error) {
	if q.Request.Method != selectMethod {
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

func (q *queryBuilder) Print() *queryBuilder {
	log.Println(q.Request)
	return q
}





