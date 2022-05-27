package http

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	pb "github.com/naufal-dean/ccb/protobuf"
)

func convertHeader(header http.Header) map[string]string {
	// Get first value
	res := make(map[string]string)
	for k, _ := range header {
		res[k] = header.Get(k)
	}
	return res
}

func convertResponse(res *http.Response) (*pb.Response, error) {
	t1 := time.Now()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	t2 := time.Now()

	res2 := &pb.Response{
		Status:        res.Status,
		StatusCode:    int32(res.StatusCode),
		Proto:         res.Proto,
		ProtoMajor:    int32(res.ProtoMajor),
		ProtoMinor:    int32(res.ProtoMinor),
		Header:        convertHeader(res.Header),
		Body:          body,
		ContentLength: res.ContentLength,
	}
	t3 := time.Now()

	log.Printf("Time cr: %d; %d\n", t2.Sub(t1).Microseconds(), t3.Sub(t2).Microseconds())

	return res2, nil
}
