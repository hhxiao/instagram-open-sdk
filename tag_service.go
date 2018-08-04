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

func (t *TagService) Top(tagName string) (res *TagResponse, err error) {
	if tagName == "" {
		err = fmt.Errorf("empty tag")
		return
	}

	var r *http.Response
	if r, err = t.client.client.Get(fmt.Sprint(BASE_URL + EXPLORE + TAGS + tagName + ENDING)); err != nil {
		return
	}

	data := struct {
		GraphQL struct {
			TagResponse TagResponse `json:"hashtag"`
		} `json:"graphql"`
	}{}

	if err = json.NewDecoder(r.Body).Decode(&data); err != nil {
		res = nil
	}

	res = &data.GraphQL.TagResponse
	res.Name = tagName

	for i := range res.TopPosts.Nodes {
		node := &res.TopPosts.Nodes[i].Node
		node.client = t.client
		if len(node.EdgeMediaToCaption.Edges) != 0 {
			node.Caption = node.EdgeMediaToCaption.Edges[0].Node.Text
		}
	}
	for i := range res.Medias.Nodes {
		node := &res.Medias.Nodes[i].Node
		node.client = t.client
		if len(node.EdgeMediaToCaption.Edges) != 0 {
			node.Caption = node.EdgeMediaToCaption.Edges[0].Node.Text
		}
	}

	return
}

func (t *TagService) Recent(tagName string) (res *TagResponse, err error) {
	if tagName == "" {
		err = fmt.Errorf("empty tag")
		return
	}

	var r *http.Response
	if r, err = t.client.client.Get(fmt.Sprint(BASE_URL + EXPLORE + TAGS + tagName + ENDING)); err != nil {
		return
	}

	data := struct {
		GraphQL struct {
			TagResponse TagResponse `json:"hashtag"`
		} `json:"graphql"`
	}{}

	if err = json.NewDecoder(r.Body).Decode(&data); err != nil {
		res = nil
	}

	res = &data.GraphQL.TagResponse
	res.Name = tagName

	for i := range res.TopPosts.Nodes {
		res.TopPosts.Nodes[i].Node.client = t.client
	}
	for i := range res.Medias.Nodes {
		res.Medias.Nodes[i].Node.client = t.client
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

	data := struct {
		GraphQL struct {
			TagResponse TagResponse `json:"hashtag"`
		} `json:"graphql"`
	}{}

	if err = json.NewDecoder(r.Body).Decode(&data); err != nil {
		res = nil
	}

	res = &data.GraphQL.TagResponse
	res.Name = tagName

	for i := range res.TopPosts.Nodes {
		res.TopPosts.Nodes[i].Node.client = t.client
	}
	for i := range res.Medias.Nodes {
		res.Medias.Nodes[i].Node.client = t.client
	}

	return
}
