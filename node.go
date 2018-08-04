package ig

import (
	"image"
	"image/jpeg"
	"net/http"
)

type Node struct {
	Caption            string       `json:"caption"`
	ShortCode          string       `json:"shortcode"`
	Comments           CountWrapper `json:"edge_media_to_comment"`
	Date               int          `json:"taken_at_timestamp"`
	Dimensions         Dimensions   `json:"dimensions"`
	DisplaySrc         string       `json:"display_url"`
	ID                 string       `json:"id"`
	IsVideo            bool         `json:"is_video"`
	Likes              CountWrapper `json:"edge_media_preview_like"`
	Owner              IdWrapper    `json:"owner"`
	ThumbnailSrc       string       `json:"thumbnail_src"`
	EdgeMediaToCaption struct {
		Edges []struct {
			Node struct {
				Text string `json:"text"`
			} `json:"node"`
		} `json:"edges"`
	} `json:"edge_media_to_caption"`
	ThumbnailResources []struct {
		Src    string `json:"src"`
		Width  int    `json:"config_width"`
		Height int    `json:"config_height"`
	} `json:"thumbnail_resources"`
	client
}

func (n *Node) GetImage() (image.Image, error) {
	return getImage(n.client.client, n.DisplaySrc)
}

func (n *Node) GetThumbnail() (image.Image, error) {
	return getImage(n.client.client, n.DisplaySrc)
}

func getImage(c *http.Client, imgUrl string) (img image.Image, err error) {
	var r *http.Response
	if r, err = c.Get(imgUrl); err != nil {
		return
	}
	img, err = jpeg.Decode(r.Body)
	r.Body.Close()
	return
}
