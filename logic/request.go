package logic

import (
	"math/rand"
	"net/http"
	"time"
)

const (
	requestInterval = time.Second * 1
)

type Requester struct {
	srvName string
	ticker  *time.Ticker
	stop    chan struct{}
}

func NewRequester(srvName string) *Requester {
	return &Requester{
		srvName: srvName,
	}
}

func (r *Requester) StartRequest() {
	r.stop = make(chan struct{})
	r.ticker = time.NewTicker(requestInterval)
	uris := MustLoadConfig().Server[r.srvName].Uri

	go func() {
		for {
			select {
			case <-r.stop:
				return
			case <-r.ticker.C:
				r.request(uris[rand.Intn(len(uris))])
			}
		}
	}()
}

func (r *Requester) StopRequest() {
	r.ticker.Stop()
	close(r.stop)
}

func (r *Requester) request(url string) error {
	_, err := http.Get(FormatRequestURL(r.srvName, url))
	if err != nil {
		panic(err)
	}
	return nil
}
