package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)



func main() {
		var wg sync.WaitGroup
		wg.Add(2)

		go startServer1(&wg)
		go startServer2(&wg)


		wg.Wait()

}

func startServer1(wg *sync.WaitGroup)  {
	defer wg.Done()
	var httpServer1 http.Server
	httpServer1.Addr = "127.0.0.1:8080"
	if err := httpServer1.ListenAndServe();err != nil {
		if err == http.ErrServerClosed {
			log.Println("HTTP server 1 closed.")
		} else {
			log.Printf("HTTP server 1 error: %v\n",err)
		}
	}
}

func startServer2(wg *sync.WaitGroup) {
	defer wg.Done()
	mux1 := http.NewServeMux()
	mux1.HandleFunc("/hi", func(w http.ResponseWriter, req *http.Request) {
		//if req.URL.Path != "hi"{
		if req.URL.Path != "/hi"{
			http.NotFound(w, req)
			return
		}
		name := req.FormValue("name")
		if  name == "" {
			fmt.Fprint(w, "Welcome!")
		} else {
			fmt.Fprintf(w, "Welcome, %s!", name)
		}
	})
	httpServer2 := http.Server{
		Addr:              "127.0.0.1:8081",
		Handler:           mux1,
	}
	if err := httpServer2.ListenAndServe();err != nil{
		if err == http.ErrServerClosed{
			log.Println("HTTP server 2 closed.")
		} else {
			log.Printf("HTTP server 2 error: %v\n", err)
		}
	}
}

//curl -X POST 127.0.0.1:8081/hi -d "name=xx"