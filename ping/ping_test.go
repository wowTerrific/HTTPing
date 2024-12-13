package ping

import (
	"testing"
)

func TestPing(t *testing.T) {

	pingTests := []struct {
		name string
		url  string
		want string
	}{
		{"return 200 OK from google.com", "https://google.com", "200 OK"},
		{"return error from bad url", "abd://abcdefg.atub", "Get \"abd://abcdefg.atub\": unsupported protocol scheme \"abd\""},
		{"return 404 Not Found from made-up url", "https://www.google.com/asfjasdlfjaslfj", "404 Not Found"},
	}

	for _, tt := range pingTests {
		t.Run(tt.name, func(t *testing.T) {
			got := Ping(tt.url)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
