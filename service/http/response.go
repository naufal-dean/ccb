package http

import (
	"io/ioutil"
	"net/http"

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
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &pb.Response{
		Status:        res.Status,
		StatusCode:    int32(res.StatusCode),
		Proto:         res.Proto,
		ProtoMajor:    int32(res.ProtoMajor),
		ProtoMinor:    int32(res.ProtoMinor),
		Header:        convertHeader(res.Header),
		Body:          body,
		ContentLength: res.ContentLength,
	}, nil
}
