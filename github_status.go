package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

type Rej struct{
	Status string `json:"status"`
	Body string `json:"body"`
	Created string `json:"created_on"`
}

func main() {
	uri := "https://status.github.com/api/last-message.json"
	req,_ := http.NewRequest("GET",uri,nil)

	client := new(http.Client)
	resp,_ := client.Do(req)
	defer resp.Body.Close()

	byteArray,_ := ioutil.ReadAll(resp.Body)
	result := string(byteArray)
	//fmt.Println(string(byteArray))
	//fmt.Println(byteArray)
	fmt.Println(result)
}
