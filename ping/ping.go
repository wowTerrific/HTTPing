package ping

import (
	"net/http"
)

func Ping(url string) string {
	res, err := http.Get(url)

	if err != nil {
		return err.Error()
	}

	return res.Status
}
