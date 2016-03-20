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
		"cdr_id" : rtu.OrderTypeAsc,
	}

	res, err := query.Select().From(rtu.TableCDRHour).Where("in_ani = 11111111111").OrderBy(sort).Limit(2).Offset(1).GetCDRs()

	if err != nil {
		log.Println(err)
	}

	for in, it := range res.Items {
		log.Printf("%v, %v", in, it)
	}

	/*Select with complex filter example

	sort := map[string]rtu.OrderType{
		"cdr_id" : rtu.OrderTypeAsc,
	}

	filters := []string {
		"in_ani=12345678901",
		"out_ani=10987654321",
	}

	res, err := query.Select().From(rtu.TableCDRHour).Filters("and", filters).OrderBy(sort).Limit(2).Offset(1).GetCDRs()

	if err != nil {
		log.Println(err)
	}

	for in, it := range res.Items {
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

	res, err := query.Count(rtu.TableCDRHour).Where("in_ani = 11111111111").GetInt()
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
	res, err := query.Insert().Into(rtu.TablePrerouting).Values(rowset).GetInt()
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

	res, err := query.Update(rtu.TablePrerouting).Set(rowset).Where("rule_name=testrule").GetInt()

	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("%v", res)
	*/

	/* Delete example


	res, err := query.Delete().From(rtu.TablePrerouting).Where("rule_name=testrule").GetInt()
	if err != nil {
		log.Println(err)
	}

	log.Printf("%v", res)
	 */
}


