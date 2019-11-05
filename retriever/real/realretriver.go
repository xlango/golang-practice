package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriver struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r *Retriver) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	rs, err := httputil.DumpResponse(resp, true)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	return string(rs)
}
