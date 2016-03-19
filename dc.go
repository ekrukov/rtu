package rtu

import (
	"errors"
)

var (
	errUnknownTable = errors.New("Table id not found in dictionary")
)
var queryDefaultLimit int = 1000

var urlMake = func(s string) string {
	return "https://" + s + "/service/service.php?soap"
}

type methodType string

const (
	selectMethod methodType = "selectRowset"
	insertMethod methodType = "insertRowset"
	updateMethod methodType = "updateRowset"
	deleteMethod methodType = "deleteRowset"
	countMethod methodType = "countRowset"
	describeMethod methodType = "describeColumns"
)

var filterConditions = []string{"<=", ">=", "<>", "like", "not like", "regexp", ">", "<", "="}

type Ordertype string

const (
	OrdertypeAsc Ordertype = "asc"
	OrdertypeDesc Ordertype = "desc"
)

var tableIds = map[string]string{
	"cdrH": "02.2205.01",
	"cdrD": "02.2206.01",
	"cdrW": "02.2207.01",
	"cdrM": "02.2208.01",
	"cdrA": "02.2204.01",
	"prerouting" : "02.2211.01",

}

func GetTableIdByName(n string) (id string, err error) {
	if id, ok := tableIds[n]; ok {
		return id, nil
	} else {
		return "", errUnknownTable
	}
}


