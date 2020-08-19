package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type httpWordRepository struct {
	client *http.Client
}

func NewHTTPWordRepository() WordRepository {
	return &httpWordRepository{
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (repo *httpWordRepository) Get(ctx context.Context, count int) ([]string, error) {
	url := fmt.Sprintf("https://random-word-api.herokuapp.com/word?number=%d", count)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := repo.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("request_failed")
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var words []string
	err = json.Unmarshal(responseBody, &words)

	return words, err

}
