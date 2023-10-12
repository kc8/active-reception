package requests

import (
	resty "github.com/go-resty/resty/v2"
    "fmt"
)


type Method string
const (
    GET = "GET"
    PUT = "PUT"
)

type RequestInfo struct {

}

type PerformRequest func(string) string

func SendGetRequset(path string) string {
	client := resty.New()
	resp, err := client.R().Get(path)
    if err != nil {
        fmt.Println(err)
    }
	return resp.String()
}

func SendPutRequest() {
}

func GetCorrectRequsetFunc(method string) PerformRequest {
    switch method {
        case GET:
            return SendGetRequset
        case PUT:
            return SendGetRequset
    }
    return nil
}
