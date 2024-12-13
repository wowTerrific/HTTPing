package ping

import (
	"fmt"
	"net/http"
)

func Ping(url string) string {
	res, err := http.Get(url)

	if err != nil {
		return fmt.Sprintf("%s - %s", url, err.Error())
	}

	return fmt.Sprintf("%s - %s", url, res.Status)
}
