package main

import (
	"log"
	"net/http"
	"sync"
)



func main() {
		var wg sync.WaitGroup
		wg.Add(2)

		go startServer1(&wg)

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