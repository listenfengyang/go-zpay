package utils

import (
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

type RestyRequest struct {
	Method   string              `json:"method"`
	Url      string              `json:"url"`
	Headers  http.Header         `json:"headers"`
	Body     interface{}         `json:"body"`
	FormData map[string][]string `json:"formData"`
}

type RestyResponse struct {
	StatusCode int         `json:"status_code"`
	Status     string      `json:"status"`
	Headers    http.Header `json:"headers"`
	Body       string      `json:"body"`
	ReceivedAt time.Time   `json:"received_at"`
	Rtt        int64       `json:"rtt"`
}

type RestyLog struct {
	Request  RestyRequest  `json:"request"`
	Response RestyResponse `json:"response"`
}

//--------------------------------------------------------

func GetRestyLog(resp *resty.Response) RestyLog {
	//request header
	reqHeaders := map[string][]string(resp.Request.Header)
	delete(reqHeaders, "User-Agent")

	return RestyLog{
		Request: RestyRequest{
			Method:   resp.Request.Method,
			Url:      resp.Request.URL,
			Headers:  reqHeaders, //resp.Request.Header,
			Body:     resp.Request.Body,
			FormData: resp.Request.FormData,
			//"time":    resp.Request.Time,
		},
		Response: RestyResponse{
			StatusCode: resp.StatusCode(),
			Status:     resp.Status(),
			Headers:    resp.Header(),
			Body:       resp.String(),
			ReceivedAt: resp.ReceivedAt(),
			Rtt:        resp.Time().Milliseconds(),
			//"time":        resp.Time(),
		},
	}
}
