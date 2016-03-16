package main

import (
	"github.com/ekrukov/rtu"
	"log"
)

var serverName string = "localhost"
var serverLogin string = "admin"
var serverPass string = "superpasswd"

func main() {
	client := rtu.NewRTUClient(serverName, serverLogin, serverPass)


	/* Select example */

	filter := map[string]string {
		"type" : "cond",
		"column" : "in_ani",
		"operator" : "=",
		"value" : "11111111111",
	}

	res, err := rtu.NewRTUQuery(client).Select().From("cdrH").Where(filter).Run()
	if err != nil {
		log.Println(err)
	}
	for in, it := range res.Select.Result.Rows {
		for _, item := range it.Items {
			log.Printf("%v, %v, %v", in, item.Key, item.Value)
		}
		log.Printf("")
	}

	/* Describe example */

	res, err = rtu.NewRTUQuery(client).Describe("cdrH").Run()
	if err != nil {
		log.Println(err)
	}

	for in, it := range res.Describe.Result.Rows {
		for _, item := range it.Items {
			log.Printf("%v, %v, %v", in, item.Key, item.Value)
		}
		log.Printf("")
	}

	/* Count example */

	filter = map[string]string {
		"type" : "cond",
		"column" : "in_ani",
		"operator" : "=",
		"value" : "11111111111",
	}

	res, err = rtu.NewRTUQuery(client).Count("cdrH", filter).Run()
	if err != nil {
		log.Println(err)
	}

	log.Printf("%v", res.Count)



	/* Insert example */

	rowset := []map[string]string {
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
	res, err = rtu.NewRTUQuery(client).Insert().Into("prerouting").Values(&rowset).Run()
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("%v", res.Insert)

	/*  Update example */

	rowset = []map[string]string {
		0: {
			"priority" : "105",
		},
	}

	filter = map[string]string {
		"type" : "cond",
		"column" : "rule_name",
		"operator" : "=",
		"value" : "testrule",
	}

	res, err = rtu.NewRTUQuery(client).Update("prerouting").Set(&rowset).Where(filter).Run()

	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("%v", res.Update)

}


