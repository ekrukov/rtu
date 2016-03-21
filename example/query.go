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


	/*Simple select example */

	sort := map[string]rtu.OrderType{
		"CDR_ID" : rtu.OrderTypeAsc,
	}

	res, err := query.Select().From(rtu.TableCDRHour).Where("IN_ANI = 11111111111").OrderBy(sort).Limit(2).Offset(1).GetCDRs()

	if err != nil {
		log.Println(err)
	}

	for in, it := range res {
		log.Printf("%v, %v", in, it)
	}

	/*Select with complex filter example

	sort := map[string]rtu.OrderType{
		"CDR_ID" : rtu.OrderTypeAsc,
	}

	filters := []string {
		"IN_ANI=12345678901",
		"ELAPSED_TIME >=1",
	}

	res, err := query.Select().From(rtu.TableCDRHour).Filters(rtu.FilterHandleAND, filters).OrderBy(sort).Limit(2).Offset(1).GetCDRs()

	if err != nil {
		log.Println(err)
	}

	for in, it := range res {
		log.Printf("%v, %v", in, it)
	}

	/* Describe example

	res, err := query.Describe(rtu.TableCDRHour).GetRows()
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

	res, err := query.Count(rtu.TableCDRHour).Where("IN_ANI = 11111111111").GetInt()
	if err != nil {
		log.Println(err)
	}

	log.Printf("%v", res)



	/* Insert example

	client := rtu.NewRTUClient(serverName, serverLogin, serverPass)

	pr := &rtu.Prerouting{
		RULE_NAME : "testrule",
		PRIORITY : "100",
		DISCONNECT_CODE : "262546",
		ACTION : "2",
		DESCRIPTION : "testdesc",
		ANI_PATTERN : "11111111111",
		DNIS_EXCLUDE : "1111111111[0-9]",
	}
	res, err := pr.Insert(client)
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

	res, err := query.Update(rtu.TablePrerouting).Set(rowset).Where("RULE_NAME=testrule").GetInt()

	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("%v", res)
	*/

	/* Delete example


	res, err := query.Delete().From(rtu.TablePrerouting).Where("RULE_NAME=testrule").GetInt()
	if err != nil {
		log.Println(err)
	}

	log.Printf("%v", res)
	 */
}


