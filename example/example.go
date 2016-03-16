package main

import (
	"github.com/ekrukov/rtu"
	"github.com/ekrukov/rtu/soap"
	"log"
)

var serverName string = "localhost"
var serverLogin string = "admin"
var serverPass string = "superpasswd"

func main() {
	client := rtu.NewRTUClient(serverName, serverLogin, serverPass)

	/* Select example

	res, err := rtu.NewRTUQuery(client).Select().From("cdrH").Where(&soap.Filter{
		Type_: "cond",
		Column: "in_ani",
		Operator: "=",
		Value: "1111111111",
	}).Run()
	if err != nil {
		log.Println(err)
	}
	for in, it := range res.Select.Result.Rows {
		for _, item := range it.Items {
			log.Printf("%v, %v, %v", in, item.Key, item.Value)
		}
		log.Printf("")
	}

	*/

	/*Describe example

	res, err := rtu.NewRTUQuery(client).Describe("cdrH").Run()
	if err != nil {
		log.Println(err)
	}

	for in, it := range res.Describe.Result.Rows {
		for _, item := range it.Items {
			log.Printf("%v, %v, %v", in, item.Key, item.Value)
		}
		log.Printf("")
	}

	*/

	/* Count example

	res, err := rtu.NewRTUQuery(client).Count("cdrH", &soap.Filter{
		Type_: "cond",
		Column: "in_ani",
		Operator: "=",
		Value: "1111111111",
	}).Run()
	if err != nil {
		log.Println(err)
	}

	log.Printf("%v", res.Count)

	*/

	/* Insert example

	rowsetMap := []map[string]string {
		0: {
			"rule_name" : "testrule",
			"priority" : "100",
			"disconnect_code" : "262546",
			"action" : "2",
			"description" : "testdesc",
			"ani_pattern" : "11111111111",
			"dnis_exclude" : "1800.{7}",
		},
	}
	rowset, err := soap.MapsToRowset(&rowsetMap)
	res, err := rtu.NewRTUQuery(client).Insert().Into("prerouting").Values(rowset).Run()
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("%v", res.Insert)

	*/

	/*  Update example

	rowsetMap := []map[string]string {
		0: {
			"priority" : "101",
		},
	}
	rowset, err := soap.MapsToRowset(&rowsetMap)
	if err != nil {
		log.Println(err)
		return
	}
	res, err := rtu.NewRTUQuery(client).Update("prerouting").Set(rowset).Where(&soap.Filter{
		Type_: "cond",
		Column: "rule_name",
		Operator: "=",
		Value: "testrule",
	}).Run()

	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("%v", res.Update)
	*/
}


