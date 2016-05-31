package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
	"github.com/deckarep/gosx-notifier"
)

// json struct
type ResponseJson struct{
	Status string `json:"status"`
	Last_updated string `json:"last_updated"`
}

// notification
func notifier(status string,created string){
	note := gosxnotifier.NewNotification(status+"\n"+created)

	note.Title = "Github status"
	//note.Subtitle = b 
	//note.Sound = gosxnotifier.Default
	note.Group = "Go_status"
	note.Link = "https://status.github.com/"
	note.AppIcon = "./Github_Mark.png"

	err := note.Push()
	if err != nil {
		log.Println("Error")
	}
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

	//fmt.Printf("%v\n%v\n",result.Status,result.Last_updated)
	notifier(result.Status,result.Last_updated)
}
