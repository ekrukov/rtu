package rtu

import (
	"github.com/ekrukov/rtu/soap"
	"log"
	"errors"
)

const (
//TODO
	OrdertypeAsc soap.Ordertype = soap.OrdertypeAsc
	OrdertypeDesc soap.Ordertype = soap.OrdertypeDesc
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
		log.Fatal("")
		return "", errors.New("Attempt to search for unknown table")
	}
}


