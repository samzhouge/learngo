package real2

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	response, err := httputil.DumpResponse(resp, true)
	defer resp.Body.Close()
	//all, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(response)
}
