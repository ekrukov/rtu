package main

import (
        "github.com/ekrukov/rtu"
	"github.com/ekrukov/rtu/soap"
	"log"
)

var serverName string = "localhost"
var serverLogin string = "admin"
var serverPass string = "superpasswd"

func main(){
	client := rtu.NewRTUClient(serverName, serverLogin, serverPass)
	res, err := rtu.NewRTUQuery(client).Select().From("cdrH").Where(&soap.Filter{
		Type_: "cond",
		Column: "in_ani",
		Operator: "=",
		Value: "12345678987",
	}).Run()
	if err != nil {
		log.Println(err)
	}
	log.Printf("%+v", res)
}

