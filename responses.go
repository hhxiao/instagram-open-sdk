package ig

type Dimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type CountWrapper struct {
	Count int `json:"count"`
}

type IdWrapper struct {
	ID string `json:"id"`
}

type PageInfo struct {
	HasNextPage     bool   `json:"has_next_page"`
	EndCursor       string `json:"end_cursor"`
	StartCursor     string `json:"start_cursor"`
	HasPreviousPage bool   `json:"has_previous_page"`
}

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
}

type TopPosts struct {
	Nodes []Node `json:"nodes"`
}

type Media struct {
	Nodes    []Node   `json:"nodes"`
	PageInfo PageInfo `json:"page_info"`
	Count    int      `json:"count"`
	Status   string   `json:"status"`
}

type TagRecent struct {
	Name            string      `json:"name"`
	ContentAdvisory interface{} `json:"content_advisory"`
	Media           Media       `json:"media"`
}

type TagTop struct {
	Name            string      `json:"name"`
	TopPosts        TopPosts    `json:"top_posts"`
	ContentAdvisory interface{} `json:"content_advisory"`
}

type TagRecentResponse struct {
	Name string `json:"name"`
	Data Media  `json:"media"`
	client
}

type TagTopResponse struct {
	Data TopPosts `json:"top_posts"`
	client
}

type LoginResponse struct {
	Status        string `json:"status"`
	Authenticated bool   `json:"authenticated"`
	User          string `json:"user"`
}
