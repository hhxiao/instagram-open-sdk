package ig

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const tagQuery = `ig_hashtag(%s) { media.after(%s, %d) {
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

func (t *TagService) Recent(tagName string) (res *TagRecentResponse, err error) {
	if tagName == "" {
		err = fmt.Errorf("Empty tag")
		return
	}

	var r *http.Response
	if r, err = t.client.client.Get(fmt.Sprint(BASE_URL + EXPLORE + TAGS + tagName + ENDING)); err != nil {
		return
	}

	data := struct {
		TagRecent `json:"tag"`
	}{}

	if err = json.NewDecoder(r.Body).Decode(&data); err != nil {
		res = nil
	}

	res = &TagRecentResponse{
		Name:   data.TagRecent.Name,
		Data:   data.TagRecent.Media,
		client: t.client,
	}
	return
}

func (t *TagRecentResponse) NextPage(amount uint) (res *TagRecentResponse, err error) {
	if !t.Data.PageInfo.HasNextPage {
		err = fmt.Errorf("No next page available")
		return
	}

	var r *http.Response
	form := url.Values{}
	form.Add("q", fmt.Sprintf(tagQuery, t.Name, t.Data.PageInfo.EndCursor, amount))

	if r, err = t.client.post(BASE_URL+QUERY, form); err != nil {
		return
	}

	data := struct {
		Media `json:"media"`
	}{}

	if err = json.NewDecoder(r.Body).Decode(&data); err != nil {
		res = nil
	}

	res = &TagRecentResponse{
		Name:   t.Name,
		Data:   data.Media,
		client: t.client,
	}
	return
}
