package ig

import (
	"fmt"
	"log"
	"testing"
)

// Set these to run tests
const (
	user     = ""
	password = ""
)

var c *Client
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

func TestTag(t *testing.T) {
	res, err := c.TagService.Recent("mitäserikkaillekuuluumitäköyhättekee")
	if err != nil {
		t.Error(err)
		return
	}

	var counter int
	for i, m := range res.Data.Nodes {
		fmt.Println(i, m.Caption[0:10])
	}
	for err == nil {
		counter += len(res.Data.Nodes)
		if res, err = res.NextPage(10); err != nil {
			t.Error(err)
			return
		}
		for i, m := range res.Data.Nodes {
			min := len(m.Caption)
			if min > 10 {
				min = 10
			}
			fmt.Println(i+counter, m.Caption[0:min])
		}
	}
}
