package ping

import (
	"log"
	"net/http"
	"os"
)

func Ping(url string) string {
	res, err := http.Get(url)

	if err != nil {
		log.Fatalf("error in ping: %v", err)
		os.Exit(1)
	}

	return res.Status
}
