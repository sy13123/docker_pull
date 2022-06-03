package cmd

import (
	//"fmt"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type req struct{
	client *resty.Request 
	url string 
	prames map[string]interface{}
	data map[string]interface{}
}


func (c req) get(url string,url1 map[string]interface{}) *resty.Response {
	rep,err := c.client.Get(url)
	fmt.Println(err)
	return rep
} 


func requests() *req {
	client := resty.New()
	r :=&req{client:client.R()} 
	return r
	//Get("https://baidu.com")
	//return resp,err
//
	//if err != nil {
		//fmt.Print(err)
	//	return err
	//  //log.Fatal(err)
	//}
    //return resp
	//fmt.Println("Response Info:",AuthSuccess)
 	//fmt.Println("Status Code:", resp.StatusCode())
	//fmt.Println("Status:", resp.Status())
	//fmt.Println("Proto:", resp.Proto())
	//fmt.Println("Time:", resp.Time())
	//fmt.Println("Received At:", resp.ReceivedAt())
	//fmt.Println("Size:", resp.Size())
	//fmt.Println("Headers:") 
	//for key, value := range resp.Header() {
	//  fmt.Println(key, "=", value)
	//}
	//fmt.Println("Cookies:")
	//for i, cookie := range resp.Cookies() {
	//  fmt.Printf("cookie%d: name:%s value:%s\n", i, cookie.Name, cookie.Value)
	//}
  }