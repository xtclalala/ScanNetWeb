package requests

import (
	"github.com/pkg/errors"
	"net/http"
	"sync"
)

type Requests struct {
	*http.Request
	*http.Client
	*http.Response
	Error error
}

func (r Requests) GetStatus() int {
	return r.Response.StatusCode
}

func (r Requests) Do(req *http.Request) *Requests {
	if req == nil {
		r.do(r.Request)
		return &r
	}
	r.do(req)
	return &r
}

func (r Requests) do(req *http.Request) {
	if r.Error == nil {
		response, err := r.Client.Do(req)
		r.Error = errors.Wrap(err, "The Request Failed")
		r.Response = response
	}
}

func NewRequests(method, url string, header map[string]string) *Requests {
	request, _ := http.NewRequest(method, url, nil)
	if header != nil {
		for key, value := range header {
			request.Header.Add(key, value)
		}
	}
	return &Requests{Request: request, Client: &http.Client{}}
}

func Get(url string, header *map[string]string) *Requests {
	req := NewRequests("GET", url, *header)
	return req.Do(req.Request)
}

func SyncGet(urls []string, header *map[string]string) *sync.Map {
	respMap := sync.Map{}
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string, header *map[string]string, respMap *sync.Map) {
			defer wg.Done()
			requests := Get(url, header)
			respMap.Store(url, requests)
		}(url, header, &respMap)
	}
	wg.Wait()
	return &respMap
}
