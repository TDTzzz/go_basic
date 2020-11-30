package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	exp5()
}

//普通的解析
func exp1() {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		Id   int    `json:"id"`
	}
	jsonData := []byte(`
    {
        "name": "tt",
        "age": 23,
        "id": 999
    }`)

	var person Person
	err := json.Unmarshal(jsonData, &person)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(person)
}

//解析内嵌对象数组的JSON
func exp2() {
	type Friend struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	type Person struct {
		Name   string   `json:"name"`
		Age    int      `json:"age"`
		Id     int      `json:"id"`
		Friend []Friend `json:"friend"`
	}

	jsonData := []byte(`
    {
        "name": "tt",
        "age": 23,
        "id": 999,
		"friend":[
			{
				"name":"hh",
				"age":24
			}
		]
    }`)
	var person Person
	err := json.Unmarshal(jsonData, &person)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(person)
}

//解析具有动态Key的对象
func exp3() {
	type Friend struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	type Person struct {
		Name   string            `json:"name"`
		Age    int               `json:"age"`
		Id     int               `json:"id"`
		Friend map[string]Friend `json:"friend"`
	}

	jsonData := []byte(`
    {
        "name": "tt",
        "age": 23,
        "id": 999,
		"friend":{
			"ss":{
				"name":"hh",
				"age":24
			}
		}
    }`)
	var person Person
	err := json.Unmarshal(jsonData, &person)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(person)
}

//解析包含任意层级的数组和对象的JSON数据
func exp4() {
	jsonData := []byte(`{"Name":"Eve","Age":6,"Parents":["Alice","Bob"]}`)

	var v interface{}
	json.Unmarshal(jsonData, &v)
	data := v.(map[string]interface{})

	//类型断言
	for k, v := range data {
		switch v := v.(type) {
		case string:
			fmt.Println(k, v, "(string)")
		case float64:
			fmt.Println(k, v, "(float64)")
		case []interface{}:
			fmt.Println(k, "(array):")
			for i, u := range v {
				fmt.Println("    ", i, u)
			}
		default:
			fmt.Println(k, v, "(unknown)")
		}
	}
	log.Println(v)
}

//用decoder解析数据流
//如果json的数据载体是打开的文件或者HTTP请求这种数据流
//不必把JSON数据读取出来再用UnMarshall,直接用Decode即可读取并解析JSON
func exp5() {
	const jsonStream = `
    {"Name": "Ed", "Text": "Knock knock."}
    {"Name": "Sam", "Text": "Who's there?"}
    {"Name": "Ed", "Text": "Go fmt."}
    {"Name": "Sam", "Text": "Go fmt who?"}
    {"Name": "Ed", "Text": "Go fmt yourself!"}
`
	type Message struct {
		Name, Text string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}
