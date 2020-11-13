package recommendations

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Reccomender struct {
	url string
}

func NewRecommender(url string) Reccomender {
	return Reccomender{url}
}

func (r Reccomender) GetUserRecommendations(userId int) ([]int, error) {
	resp, err := http.Get(fmt.Sprintf("%s/recommender/%d", r.url, userId))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("%d", resp.StatusCode))
	}
	var ids []int
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &ids)
	if err != nil {
		return nil, err
	}
	return ids, nil
}
