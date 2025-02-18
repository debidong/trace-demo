package requester

import (
	"net/http"
	"strings"
	"time"
	"trace-demo/server"
)

const (
	requestInterval = time.Second * 1
)

type Requester struct {
	country string
	ticker  *time.Ticker
	stop    chan struct{}
}

func NewRequester(country string) *Requester {
	return &Requester{
		country: country,
	}
}

func (r *Requester) StartRequest() {
	r.stop = make(chan struct{})
	r.ticker = time.NewTicker(requestInterval)
	go func() {
		for {
			select {
			case <-r.stop:
				return
			case <-r.ticker.C:
				r.request(r.country)
			}
		}
	}()
}

func (r *Requester) StopRequest() {
	close(r.stop)
}

func (r *Requester) request(url string) error {
	_, err := http.Get(formatRequestURL(url))
	if err != nil {
		panic(err)
	}
	return nil
}

func formatRequestURL(url string) string {
	prefix := "http:/"
	return strings.Join([]string{prefix, server.ServerAddr, url}, "/")
}
