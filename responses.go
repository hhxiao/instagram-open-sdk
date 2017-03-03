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

type TopPosts struct {
	Nodes []Node `json:"nodes"`
}

type Media struct {
	Nodes    []Node   `json:"nodes"`
	PageInfo PageInfo `json:"page_info"`
	Count    int      `json:"count"`
	Status   string   `json:"status"`
}

type TagResponse struct {
	Name string `json:"name"`
	Data Media  `json:"media"`
}

type TopResponse struct {
	Name string   `json:"name"`
	Data TopPosts `json:"top_posts"`
}

type LoginResponse struct {
	Status        string `json:"status"`
	Authenticated bool   `json:"authenticated"`
	User          string `json:"user"`
}
