package barrier

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var timeoutMilliseconds int = 5000

type barrierResp struct {
	Err  error
	Resp string
}

func makeRequest(out chan<- barrierResp, url string) {
	res := barrierResp{}
	client := http.Client{
		Timeout: time.Duration(time.Duration(timeoutMilliseconds) * time.Millisecond),
	}

	resp, err := client.Get(url)
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	res.Resp = string(byt)
	out <- res
}

func barrier(endpoints ...string) {
	numRequests := len(endpoints)
	in := make(chan barrierResp, numRequests)
	defer close(in)

	responses := make([]barrierResp, numRequests)

	for _, endpoint := range endpoints {
		go makeRequest(in, endpoint)
	}

	var hasError bool

	for index := 0; index < numRequests; index++ {
		resp := <-in
		if resp.Err != nil {
			fmt.Println("ERROR: ", resp.Err)
			hasError = true
		}
		responses[index] = resp
	}

	if !hasError {
		for _, resp := range responses {
			fmt.Println(resp.Resp)
		}
	}
}
