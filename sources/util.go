package sources

import (
	"io"
	"net/url"

	"github.com/khulnasoft-labs/retryablehttp-go"
)

func NewHTTPRequest(method, url string, body io.Reader) (*retryablehttp.Request, error) {
	request, err := retryablehttp.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", "Reveal - FOSS Project (github.com/khulnasoft-labs/reveal)")
	return request, nil
}

func GetHostname(u string) (string, error) {
	parsedURL, err := url.Parse(u)
	if err != nil {
		return "", err
	}
	return parsedURL.Hostname(), nil
}
