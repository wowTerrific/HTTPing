package readfile

import (
	"reflect"
	"testing"
)

func TestParseConfig(t *testing.T) {
	data := []byte(`https://www.google.com/
https://facebook.com/
https://youtube.com/`)
	got, err := ParseConfig(data)
	want := []string{
		"https://www.google.com/",
		"https://facebook.com/",
		"https://youtube.com/",
	}

	if err != nil {
		t.Errorf("did not want error, got %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

func TestValidateConfig(t *testing.T) {
	data := []string{
		"https://www.google.com/",
		"https://facebook.com/",
		"https://youtube.com/",
	}
	got, err := ValidateConfig(data)

	if err != nil {
		t.Errorf("should not have recieved error, got %v", err)
	}

	if !got {
		t.Errorf("did not recieve ok from config validation")
	}

}
