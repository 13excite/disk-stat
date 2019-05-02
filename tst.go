package main

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type TeSt struct {
	K string
	V string
}
type Zzzz struct {
	My []TeSt
}

func main() {
	var my_json string = `
	[
  		{
    			"key1": 11111,
    			"key112": "val111112222"
  		},
  		{
    			"key1": 2,
    			"key2": "val2"
  		}
	]`
	var v []interface{}
	if err := json.Unmarshal([]byte(my_json), &v); err != nil {
		fmt.Println(err)
	}
	var tStr TeSt

	fmt.Println(v)
	for _, val := range v {
		fmt.Println(val)
		mapstructure.Decode(val, &tStr)
		fmt.Println(tStr)
	}

	fmt.Println(tStr)
	/*
			type Person struct {
		    Firstname string
		    Lastname  string
		    Address   struct {
		        City  string
		        State string
		    }
		}

		func main() {
		    mapPerson := make(map[string]interface{})
		    var person Person
		    mapPerson["firstname"] = "Nic"
		    mapPerson["lastname"] = "Raboy"
		    mapAddress := make(map[string]interface{})
		    mapAddress["city"] = "San Francisco"
		    mapAddress["state"] = "California"
		    mapPerson["address"] = mapAddress
		    mapstructure.Decode(mapPerson, &person)
		    fmt.Println(person)
		}
	*/
}
