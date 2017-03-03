package ig

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	tagQuery = `ig_hashtag(%s) { media.after(%s, %d) {
  count,
  nodes {
    caption,
    code,
    comments {
      count
    },
    comments_disabled,
    date,
    dimensions {
      height,
      width
    },
    display_src,
    id,
    is_video,
    likes {
      count
    },
    owner {
      id
    },
    thumbnail_src,
    video_views
  },
  page_info
}}`
)

type TagService struct {
	client
}

func (t *TagService) Top(tagName string) (res *TopResponse, err error) {
	if tagName == "" {
		err = fmt.Errorf("Empty tag")
		return
	}

	var r *http.Response
	if r, err = t.client.client.Get(fmt.Sprint(BASE_URL + EXPLORE + TAGS + tagName + ENDING)); err != nil {
		return
	}

	data := struct {
		TopResponse `json:"tag"`
	}{}

	if err = json.NewDecoder(r.Body).Decode(&data); err != nil {
		res = nil
	}

	res = &data.TopResponse

	for i, _ := range res.Data.Nodes {
		res.Data.Nodes[i].client = t.client
	}

	return
}

func (t *TagService) Recent(tagName string) (res *TagResponse, err error) {
	if tagName == "" {
		err = fmt.Errorf("Empty tag")
		return
	}

	var r *http.Response
	if r, err = t.client.client.Get(fmt.Sprint(BASE_URL + EXPLORE + TAGS + tagName + ENDING)); err != nil {
		return
	}

	data := struct {
		TagResponse `json:"tag"`
	}{}

	if err = json.NewDecoder(r.Body).Decode(&data); err != nil {
		res = nil
	}

	res = &data.TagResponse

	for i, _ := range res.Data.Nodes {
		res.Data.Nodes[i].client = t.client
	}

	return
}

func (t *TagService) After(tagName, cursor string, amount uint) (res *TagResponse, err error) {
	var r *http.Response
	form := url.Values{}
	form.Add("q", fmt.Sprintf(tagQuery, tagName, cursor, amount))

	if r, err = t.client.post(BASE_URL+QUERY, form); err != nil {
		return
	}

	res = &TagResponse{}
	if err = json.NewDecoder(r.Body).Decode(res); err != nil {
		res = nil
	}

	res.Name = tagName

	for i, _ := range res.Data.Nodes {
		res.Data.Nodes[i].client = t.client
	}
	return
}
