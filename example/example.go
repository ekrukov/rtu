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


	rowmap := map[string]string{
		"rule_name" : "testrule",
		"priority" : "100",
		"disconnect_code" : "262546",
		"action" : "2",
		"description" : "testdesc",
		"ani_pattern" : "11111111111",
		"dnis_exclude" : "1800.{7}",
	}

	row, err := soap.MapToRow(&rowmap)

	rowset := soap.Rowset{Rows: []soap.Row{*row}}
	res, err := rtu.NewRTUQuery(client).Insert().Into("prerouting").Values(&rowset).Run()
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("%v", res.Insert)

	*/
}


