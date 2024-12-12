package ping

import "testing"

func TestPing(t *testing.T) {
	

	got := Ping("https://google.com")
	want := "200 OK"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
