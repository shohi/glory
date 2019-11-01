package httputil

import (
	"io"
	"io/ioutil"
	"net/http"
)

func ReadBodyAndClose(r *http.Response) ([]byte, error) {
	if r.Body == nil {
		return nil, nil
	}
	defer r.Body.Close()

	return ioutil.ReadAll(r.Body)
}

func DiscardBodyAndClose(r *http.Response) error {
	if r.Body == nil {
		return nil
	}
	defer r.Body.Close()

	io.Copy(ioutil.Discard, r.Body)

	return nil
}
