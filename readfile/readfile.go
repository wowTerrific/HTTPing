package readfile

import "strings"

// Read entire file contents into memory with os.ReadFile

func ParseConfig(rawData []byte) ([]string, error) {
	splitData := strings.Split(strings.Replace(string(rawData), "\r\n", "\n", -1), "\n")
	return splitData, nil
}

// need to return line where error exists too
func ValidateConfig(data []string) (bool, error) {
	return false, nil
}
