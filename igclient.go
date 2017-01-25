package ig

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

const (
	BASE_URL = "https://www.instagram.com"
	LOGIN    = "/accounts/login/ajax/"
	EXPLORE  = "/explore"
	TAGS     = "/tags/"
	QUERY    = "/query/"
	ENDING   = "/?__a=1"
)

type Client struct {
	username   string
	password   string
	TagService *TagService
	client
}

type TagService struct {
	client
}

type client struct {
	client *http.Client
}

func NewClient(username, password string) (*Client, error) {
	jar, _ := cookiejar.New(nil)
	c := &http.Client{
		Jar: jar,
	}
	conn := client{c}
	client := &Client{
		username:   username,
		password:   password,
		client:     conn,
		TagService: &TagService{conn},
	}

	if err := client.login(); err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Client) login() (err error) {
	// Get login page to get cookies and csrf token
	var r *http.Response
	if _, err = c.client.client.Get(BASE_URL + LOGIN); err != nil {
		return
	}

	form := url.Values{}
	form.Add("username", c.username)
	form.Add("password", c.password)

	if r, err = c.client.post(BASE_URL+LOGIN, form); err != nil {
		return
	}

	res := &LoginResponse{}
	if err = json.NewDecoder(r.Body).Decode(res); err != nil {
		return
	}

	if res.Status != "ok" || !res.Authenticated {
		err = fmt.Errorf("Login failed")
	}
	return
}

func (c *client) post(u string, form url.Values) (r *http.Response, err error) {
	var req *http.Request
	if req, err = http.NewRequest("POST", u, bytes.NewBufferString(form.Encode())); err != nil {
		return
	}

	iu, _ := url.Parse(BASE_URL)
	cookies := c.client.Jar.Cookies(iu)
	var token string
	for _, cookie := range cookies {
		if cookie.Name == "csrftoken" {
			token = cookie.Value
		}
	}

	if token == "" {
		err = fmt.Errorf("Could not get csrf token")
		return
	}

	req.Header.Add("referer", BASE_URL)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("x-csrftoken", token)

	return c.client.Do(req)
}

func bodyPrint(in io.ReadCloser) (out io.ReadCloser) {
	b, _ := ioutil.ReadAll(in)
	fmt.Println(string(b))
	return ioutil.NopCloser(bytes.NewReader(b))
}
