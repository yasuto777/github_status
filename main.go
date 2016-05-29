package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

//type Rej struct{
//	Status string `json:"status"`
//	Body string `json:"body"`
//	Created string `json:"created_on"`
//}

type ResponseJson struct{
	Status string `json:"status"`
	Last_updated string `json:"last_updated"`
}

func main() {
	//uri := "https://status.github.com/api/last-message.json"
	uri := "https://status.github.com/api/status.json"
	req,_ := http.NewRequest("GET",uri,nil)

	client := new(http.Client)
	resp,_ := client.Do(req)
	defer resp.Body.Close()

	//[]byte
	byteArray,_ := ioutil.ReadAll(resp.Body)
	var result ResponseJson
	err := json.Unmarshal(byteArray,&result)
	if err != nil{
		fmt.Println("error:",err)
		return
	}

	fmt.Printf("%v\n%v\n",result.Status,result.Last_updated)
	//fmt.Printf("%v\n%v\n%v\n",result.Status,result.Body,result.Created)

}

