package main

import (
	"fmt"
	"net/http"
)

func main() {
	host := "google.cn"
	
	url1 := "http://" + host
	fmt.Printf("Send request to %q with method GET ...\n", url1)
	resp1, err := http.Get(url1)
	if err != nil {
		fmt.Printf("request sending error:%v\n", err)
		return
	}
	defer resp1.Body.Close()
	line1 := resp1.Proto + " " + resp1.Status
	fmt.Printf("The firest line of response:\n%s\n", line1)
	fmt.Println("",)

}