package rtu

import (
	"log"
	"errors"
)

var (
	errNeedTable = errors.New("RTU Query: need table for request")
	errNeedFilter = errors.New("RTU Query: need filter for request")
	errNeedRowset = errors.New("RTU Query: need rowset for request")
	errSetWOUpdate = errors.New("RTU Query: attempt to Set without Update")
	errFromWOSelectOrDelete = errors.New("RTU Query: attempt to From without Select or Delete")
	errIntoWOInsert = errors.New("RTU Query: attempt to Into action without Insert")
	errMethodNotFound = errors.New("RTU Query: method not found")
	errMethodUnsupport = errors.New("RTU Query: method not support in this context")
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

func (q *queryBuilder) Update(table TableName) *queryBuilder {
	q.Request.Method = updateMethod
	q.Request.Table.P_table_hi = table
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
		q.err = errSetWOUpdate
	} else {
		q.Request.Rowset, q.err = mapsToRowset(rowset)
	}
	return q
}

func (q *queryBuilder) From(table TableName) *queryBuilder {
	if q.Request.Method == selectMethod || q.Request.Method == deleteMethod {
		q.Request.Table.P_table_hi = table
	} else {
		q.err = errFromWOSelectOrDelete
	}
	return q
}

func (q *queryBuilder) Into(table TableName) *queryBuilder {
	if q.Request.Method == insertMethod {
		q.Request.Table.P_table_hi = table
	} else {
		q.err = errIntoWOInsert
	}
	return q
}

func (q *queryBuilder) Values(rowset []map[string]string) *queryBuilder {
	q.Request.Rowset, q.err = mapsToRowset(rowset)
	return q
}

func (q *queryBuilder) Filters(hc FilterHandleCondition, filters []string) *queryBuilder {
	q.Request.Filter = new(requestFilter)
	q.Request.Filter.Item = &requestFilterItem{
		Type_: complexFilter,
		Operator: string(hc),
		Childs: &requestFilterChildsArr{},
	}
	q.Request.Filter.Item.Childs.Items, q.err = sliceToChildFilters(filters)
	return q
}

func (q *queryBuilder) Where(f string) *queryBuilder {
	q.Request.Filter = new(requestFilter)
	q.Request.Filter.Item, q.err = stringToFilter(f)
	return q
}

func (q *queryBuilder) OrderBy(sort map[string]OrderType) *queryBuilder {
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

func (q *queryBuilder) Describe(table TableName) *queryBuilder {
	q.Request.Method = describeMethod
	q.Request.Table.P_table_hi = table
	return q
}

func (q *queryBuilder) Count(table TableName) *queryBuilder {
	q.Request.Method = countMethod
	q.Request.Table.P_table_hi = table
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
			return errNeedTable
		}
		if q.Request.Filter != nil {
			request.requestFilter = *q.Request.Filter
		} else {
			return errNeedFilter
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
			return errNeedTable
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
			return errNeedFilter
		}
		if q.Request.Table != nil {
			request.requestTable = *q.Request.Table
		} else {
			return errNeedTable
		}
		err := q.Client.Call(q.Request.Method, &request, &q.Result.Count)
		if err != nil {
			return err
		}
	case deleteMethod:
		request := new(deleteRowsetRequest)
		if q.Request.Filter != nil {
			request.requestFilter = *q.Request.Filter
		} else {
			return errNeedFilter
		}
		if q.Request.Table != nil {
			request.requestTable = *q.Request.Table
		} else {
			return errNeedTable
		}
		err := q.Client.Call(q.Request.Method, &request, &q.Result.Delete)
		if err != nil {
			return err
		}
	case insertMethod:
		request := new(insertRowsetRequest)
		if q.Request.Rowset != nil {
			request.requestRowset = *q.Request.Rowset
		} else {
			return errNeedRowset
		}
		if q.Request.Table != nil {
			request.requestTable = *q.Request.Table
		} else {
			return errNeedTable
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
			return errNeedRowset
		}
		if q.Request.Table != nil {
			request.requestTable = *q.Request.Table
		} else {
			return errNeedTable
		}
		if q.Request.Filter != nil {
			request.requestFilter = *q.Request.Filter
		} else {
			return errNeedFilter
		}
		err := q.Client.Call(q.Request.Method, &request, &q.Result.Update)
		if err != nil {
			return err
		}
	default:
		return errMethodNotFound
	}
	return nil
}

func (q *queryBuilder) GetRaw() (*rawResult, error) {
	err := q.queryExec()
	if err != nil {
		return nil, err
	}
	return q.Result, nil
}

func (q *queryBuilder) GetRows() ([]*responseRow, error) {
	err := q.queryExec()
	if err != nil {
		return nil, err
	}
	switch q.Request.Method {
	case selectMethod:
		return q.Result.Select.Result.Items, nil
	case describeMethod:
		return q.Result.Describe.Result.Items, nil
	}
	return nil, errMethodUnsupport
}

func (q *queryBuilder) GetInt() (int, error) {
	err := q.queryExec()
	if err != nil {
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
	return 0, errMethodUnsupport
}

func (q *queryBuilder) GetCDRs() (cs []*CDR, err error) {
	if q.Request.Method != selectMethod {
		err = errMethodUnsupport
		return nil, err
	}
	cs = []*CDR{}
	rows, err := q.GetRows()
	if err != nil {
		return nil, err
	}
	for _, it := range rows {
		cdr := new(CDR)
		fillStruct(cdr, it)
		cs = append(cs, cdr)
	}
	return cs, nil
}

func (q *queryBuilder) Print() *queryBuilder {
	log.Println(q.Request)
	return q
}





