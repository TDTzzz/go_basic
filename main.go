package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//测试限流
	data := make(map[string]interface{})
	data["account"] = "tdtzzz"
	data["password"] = "123456"
	bytesData, _ := json.Marshal(data)
	res, _ := http.Post("http://127.0.0.1:8887/login", "", bytes.NewReader(bytesData))
	body, _ := ioutil.ReadAll(res.Body)

	m, _ := jsonToMap(string(body))
	token := m["token"]

	client := &http.Client{}
	url := "http://127.0.0.1:8887/sum?a=1&b=4"

	for i := 0; i < 100; i++ {
		request, err := http.NewRequest("GET", url, nil)
		request.Header.Add("Authorization", token)

		if err != nil {
			log.Println("---")
			panic(err)
		}
		response, _ := client.Do(request)
		defer response.Body.Close()
		body2, _ := ioutil.ReadAll(response.Body)
		log.Println(string(body2))
	}

}

func jsonToMap(jsonStr string) (map[string]string, error) {

	m := make(map[string]string)
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Printf("Unmarshal with error: %+v\n", err)
		return nil, err
	}

	//for k, v := range m {
	//fmt.Printf("%v: %v\n", k, v)
	//}
	return m, nil
}
