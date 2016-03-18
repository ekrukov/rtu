package main

import (
	"github.com/ekrukov/rtu"
	"log"
)

var serverName string = "localhost"
var serverLogin string = "admin"
var serverPass string = "superpasswd"

func main() {

	query := rtu.NewRTUClient(serverName, serverLogin, serverPass).Query()


	/* Select example */

	filter := map[string]string{
		"type" : "cond",
		"column" : "in_ani",
		"operator" : "=",
		"value" : "11111111111",
	}

	sort := map[string]rtu.Ordertype{
		"cdr_id" : rtu.OrdertypeAsc,
	}

	res, err := query.Select().From("cdrH").Where(filter).OrderBy(sort).Limit(2).Offset(1).GetCDRs()

	if err != nil {
		log.Println(err)
	}

	log.Printf("%+v", res)

	/* Describe example

	res, err := query.Describe("cdrH").GetRows()
	if err != nil {
		log.Println(err)
	}

	for in, it := range res {
		for _, item := range it.Items {
			log.Printf("%v, %v, %v", in, item.Key, item.Value)
		}
		log.Printf("")
	}

	/* Count example

	filter := map[string]string{
		"type" : "cond",
		"column" : "in_ani",
		"operator" : "=",
		"value" : "11111111111",
	}

	res, err := query.Count("cdrH", filter).GetInt()
	if err != nil {
		log.Println(err)
	}

	log.Printf("%v", res)



	/* Insert example

	rowset := []map[string]string{
		0: {
			"rule_name" : "testrule",
			"priority" : "100",
			"disconnect_code" : "262546",
			"action" : "2",
			"description" : "testdesc",
			"ani_pattern" : "11111111111",
			"dnis_exclude" : "1111111111[0-9]",
		},
	}
	res, err := query.Insert().Into("prerouting").Values(rowset).GetInt()
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("%v", res)

	/*  Update example

	rowset := []map[string]string{
		0: {
			"priority" : "105",
		},
	}

	filter := map[string]string{
		"type" : "cond",
		"column" : "rule_name",
		"operator" : "=",
		"value" : "testrule",
	}

	res, err := query.Update("prerouting").Set(rowset).Where(filter).GetInt()

	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("%v", res)
	*/
}


