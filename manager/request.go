package manager

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func Get(rawUrl string) ([]byte, error) {
	resp, err := http.Get(rawUrl)
	if err != nil || resp == nil || resp.StatusCode != 200 {
		// FIXME: handle error
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func GetWithParams(rawUrl string, params url.Values) ([]byte, error) {
    structUrl, err := url.Parse(rawUrl)
    if err != nil {
		// FIXME: handle error
        return nil, err
    }
    structUrl.RawQuery = params.Encode()
	resp, err := http.Get(structUrl.String())
	if err != nil || resp == nil || resp.StatusCode != 200 {
		// FIXME: handle error
		return nil, err
	}
	defer resp.Body.Close()
    return ioutil.ReadAll(resp.Body)
}

func PostForm(url string, data url.Values) error {
	resp, err := http.PostForm(url, data)
	if err != nil || resp == nil || resp.StatusCode != 200 {
		// FIXME: handle error
		return err
	}
	resp.Body.Close()
	return nil
}
