package recommendations

import (
	"encoding/json"
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
	resp, err := http.Get(fmt.Sprintf("%s/%d", r.url, userId))
	if err != nil {
		return nil, err
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
