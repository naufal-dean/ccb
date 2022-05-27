package http

import (
	"bytes"
	"log"
	"net/http"
	"time"
)

func Test() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/endpoint-rt-b1", bytes.NewBuffer(nil))
	if err != nil {
		log.Printf("Failed to create new request: %v\n", err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to create new request: %v\n", err)
		return
	}

	t1 := time.Now()
	res2, err := convertResponse(res)
	t2 := time.Now()

	duration := t2.Sub(t1).Microseconds()
	if duration > 40000 {
		log.Println(res)
		log.Println(res2)
	}
	log.Printf("Time: %d\n", duration)
}
