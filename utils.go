package hubitat

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (c *Config) buildURL(path string) *url.URL {
	pathPrefix := "apps/api/1"

	urlPath := fmt.Sprintf("%s/%s", pathPrefix, path)

	dataUrl := c.URL
	dataUrl.Path = urlPath

	params := url.Values{}
	params.Add("access_token", c.AccessToken)

	dataUrl.RawQuery = params.Encode()

	return dataUrl
}

func GetURL(url *url.URL) ([]byte, error) {
	resp, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.Body == nil {
		//return nil, ErrNoRespBody
		return nil, nil
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}
