package ig

import (
	"image"
	"image/jpeg"
	"net/http"
)

type Node struct {
	Caption      string       `json:"caption"`
	Code         string       `json:"code"`
	Comments     CountWrapper `json:"comments"`
	Date         int          `json:"date"`
	Dimensions   Dimensions   `json:"dimensions"`
	DisplaySrc   string       `json:"display_src"`
	ID           string       `json:"id"`
	IsVideo      bool         `json:"is_video"`
	Likes        CountWrapper `json:"likes"`
	Owner        IdWrapper    `json:"owner"`
	ThumbnailSrc string       `json:"thumbnail_src"`
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
