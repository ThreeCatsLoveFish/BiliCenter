package manager

import (
	"net/http"
	"strings"
)

func Post(url string, data string) error {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		// FIXME: handle error
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil || resp == nil || resp.StatusCode != 200 {
		// FIXME: handle error
		return err
	}
	resp.Body.Close()
	return nil
}
