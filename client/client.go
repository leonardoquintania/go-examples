package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	requestId := make(chan int)
	concurrency := 5
	for i := 1; i <= concurrency; i++ {
		go worker(requestId, i)
	}
	for i := 1; i <= 10; i++ {
		requestId <- i
	}
}

func worker(requestId chan int, worker int) {
	for r := range requestId {
		res, err := http.Get("http://localhost:8585/product")
		if err != nil {
			log.Fatal("Erro")
		}
		defer res.Body.Close()
		content, _ := ioutil.ReadAll(res.Body)
		fmt.Printf("Worker %d. Request Id: %d. Content: %s", worker, r, string(content))
		time.Sleep(time.Second * 2)
	}

}
