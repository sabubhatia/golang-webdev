package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type statusCode struct {
	Code    int    `json:"Code"`
	Descrip string `json:"Descrip"`
}

type statusCodes []statusCode

func main() {
	var rcvd = `[{"Code":200,"Descrip":"OK"},{"Code":301,"Descrip":"Moved permanently"},{"Code":302,"Descrip":"StatusFound"},{"Code":303,"Descrip":"StatusSeeOther"},{"Code":307,"Descrip":"StatusTemporaryRedirect"},{"Code":400,"Descrip":"StatusBadRequest"},{"Code":401,"Descrip":"StatusUnauthorized"},{"Code":402,"Descrip":"StatusPaymentRequired"},{"Code":403,"Descrip":"StatusForbidden"},{"Code":404,"Descrip":"StatusNotFound"},{"Code":405,"Descrip":"StatusMethodNotAllowed"},{"Code":418,"Descrip":"StatusTeapot"},{"Code":500,"Descrip":"StatusInternalServerError"}]`
	var data statusCodes

	if err := json.Unmarshal([]byte(rcvd), &data); err != nil {
		log.Fatal(err)
	}

	for i, d := range data {
		fmt.Printf("%d. %d: %s\n", i, d.Code, d.Descrip)
	}

	bs, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Json is: ")
	fmt.Println(string(bs))
}
