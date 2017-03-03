package ig

import (
	"image"
	"image/jpeg"
	"log"
	"os"
	"testing"
)

// Set these to run tests
const (
	user     = ""
	password = ""
	tagName  = "nofilter"
)

var c *Client
var n *Node
var cursor string
var err error

func init() {
	if user == "" || password == "" {
		log.Fatal("Set username and password to run tests")
	}
}

func TestClient(t *testing.T) {
	if c, err = NewClient(user, password); err != nil {
		t.Error(err)
	}
}

func TestRecent(t *testing.T) {
	res, err := c.TagService.Recent(tagName)
	if err != nil {
		t.Error(err)
		return
	}
	n = &res.Data.Nodes[0]
	cursor = res.Data.PageInfo.EndCursor
}

func TestAfter(t *testing.T) {
	res, err := c.TagService.After(tagName, cursor, 500)
	if err != nil {
		t.Error(err)
	}
	prev := res.Data.Nodes[0].Date
	for _, n := range res.Data.Nodes {
		if n.Date > prev {
			panic("asd")
		}
		prev = n.Date
	}
}

func TestTop(t *testing.T) {
	_, err := c.TagService.Top(tagName)
	if err != nil {
		t.Error(err)
	}
}

func TestNode(t *testing.T) {
	var img image.Image
	if img, err = n.GetImage(); err != nil {
		t.Error(err)
	}

	f, _ := os.Create("image.jpg")
	defer f.Close()
	jpeg.Encode(f, img, nil)
}
