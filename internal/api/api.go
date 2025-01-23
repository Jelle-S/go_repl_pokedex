package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Jelle-S/pokedexcli/internal/pokecache"
)

func GetAndUnmarshal[T any](url string, c *pokecache.Cache) (T, error) {
	var _default T
	var err error

	body, ok := c.Get(url)
	if !ok {
		res, err := http.Get(url)

		if err != nil {
			return _default, err
		}
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return _default, err
		}
		c.Add(url, body)
	}

	var t T
	err = json.Unmarshal(body, &t)

	if err != nil {
		return _default, err
	}

	return t, nil
}
