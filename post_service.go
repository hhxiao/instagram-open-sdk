package ig

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type PostService struct {
	client
}

func (t *PostService) Get(shortCode string) (res *PostResponse, err error) {
	if shortCode == "" {
		err = fmt.Errorf("empty shortCode")
		return
	}

	var r *http.Response
	if r, err = t.client.client.Get(fmt.Sprint(BASE_URL + POSTS + shortCode + ENDING)); err != nil {
		return
	}

	data := struct {
		GraphQL struct {
			PostResponse PostResponse `json:"shortcode_media"`
		} `json:"graphql"`
	}{}

	if err = json.NewDecoder(r.Body).Decode(&data); err != nil {
		res = nil
	}

	res = &data.GraphQL.PostResponse

	return
}
