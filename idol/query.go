package idol

import (
	"io/ioutil"
	"net/http"
)

func Query(uri string) ([]byte, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}
