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


}


