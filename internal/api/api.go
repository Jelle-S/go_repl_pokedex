package api

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetAndUnmarshal[T any](url string) (T, error) {
	var _default T
	res, err := http.Get(url)
	if err != nil {
		return _default, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return _default, err
	}

	var t T
	err = json.Unmarshal(body, &t)

	if err != nil {
		return _default, err
	}

	return t, nil
}
