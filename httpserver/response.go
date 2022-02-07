package httpserver

import (
	"io/ioutil"
	"net/http"

	pb "github.com/naufal-dean/ccb/protobuf"
)

func readBody(res *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func convertResponse(res *http.Response, body []byte) *pb.Response {
	return &pb.Response{
		Status:        res.Status,
		StatusCode:    int32(res.StatusCode),
		Proto:         res.Proto,
		ProtoMajor:    int32(res.ProtoMajor),
		ProtoMinor:    int32(res.ProtoMinor),
		Header:        convertHeader(res.Header),
		Body:          body,
		ContentLength: res.ContentLength,
	}
}

func convertHeader(header map[string][]string) map[string]*pb.ListString {
	res := make(map[string]*pb.ListString)
	for k, v := range header {
		res[k] = &pb.ListString{Values: v}
	}
	return res
}
